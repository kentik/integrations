package baseserver

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/kentik/eggs/pkg/concurrent"
	"github.com/kentik/eggs/pkg/version"

	"github.com/kentik/eggs/pkg/features"
	"github.com/kentik/eggs/pkg/properties"

	"github.com/kentik/common/cmetrics"
	"github.com/kentik/common/healthcheck"
	"github.com/kentik/golog/logger"
)

const (
	FAILURE_CODE                 = -10
	ENV_CH_NUM_CPU               = "CH_NUM_CPU"
	readinessWaitGroupContextKey = "_baseserver_ready_wg"
	subContextNameContextKey     = "_baseserver_subctx"
)

type BaseServerConfiguration struct {
	// base service properties
	ServiceName string
	VersionInfo version.VersionInfo

	// operational
	ShutdownSettleTime time.Duration

	// logging
	LogToStdout bool
	LogLevel    string
	LogPrefix   string

	// metrics
	MetricsPrefix      string
	MetricsDestination string

	// meta server properties
	MetaListen string

	// healthchecks
	HealthCheckStartupDelay time.Duration
	HealthCheckPeriod       time.Duration
	HealthCheckTimeout      time.Duration

	// legacy healthcheck
	LegacyHealthCheckListen string

	// for tests
	SkipParseFlags bool

	// props
	PropsRefreshPeriod time.Duration
}

var BaseServerConfigurationDefaults = BaseServerConfiguration{
	LogToStdout:             true,
	LogLevel:                "debug",
	MetricsDestination:      "syslog",
	MetaListen:              "localhost:0",
	LegacyHealthCheckListen: "",
	ShutdownSettleTime:      5 * time.Second,
	HealthCheckStartupDelay: 5 * time.Second,
	HealthCheckPeriod:       30 * time.Second,
	HealthCheckTimeout:      5 * time.Second,
	PropsRefreshPeriod:      5 * time.Minute,
}

type BaseServer struct {
	*BaseServerConfiguration
	hce             *HealthCheckExecutor
	metaServer      *MetaServer
	Logger          *logger.Logger
	ctx             context.Context
	cancel          context.CancelFunc
	waitGroup       sync.WaitGroup
	propertyService properties.PropertyService
	featureService  features.FeatureService
}

// Perform baseserver initialization steps -- hopefully 9 out of 10 services can just call this and Run()
func Boilerplate(serviceName string, versionInfo version.VersionInfo, defaultPropertyBacking properties.PropertyBacking) *BaseServer {
	bs := NewBaseServer(serviceName, versionInfo, "chf", defaultPropertyBacking)
	bs.ParseFlags()
	bs.Init()
	setGlobalBaseServer(bs)
	return bs
}

func NewBaseServer(serviceName string, version version.VersionInfo, metricsPrefix string, defaultPropertyBacking properties.PropertyBacking) *BaseServer {
	conf := BaseServerConfigurationDefaults
	conf.ServiceName = serviceName
	conf.VersionInfo = version
	conf.MetricsPrefix = metricsPrefix
	conf.LogPrefix = serviceName + " "

	props := properties.NewPropertyService(
		properties.NewFileSystemPropertyBacking("/props"), // highest prio: dynamic FS props
		properties.NewEnvPropertyBacking(),                // env variables can override static defaults
		defaultPropertyBacking,                            // lowest prio: static default values
	)

	bs := &BaseServer{
		BaseServerConfiguration: &conf,
		propertyService:         props,
		featureService:          features.NewFeatureService(props),
	}
	bs.waitGroup.Add(1)
	return bs
}

func (bs *BaseServer) GetPropertyService() properties.PropertyService {
	return bs.propertyService
}

func (bs *BaseServer) GetFeatureService() features.FeatureService {
	return bs.featureService
}

// Parse standard golang command line flags (both those defined by the service and baseserver).
// Exits if the resulting configuration is broken or eg. if -v is specified
func (bs *BaseServer) ParseFlags() {
	if bs.SkipParseFlags {
		return
	}

	flag_version := flag.CommandLine.Bool("v", false, "Show version and build information")

	flag.CommandLine.StringVar(&bs.ServiceName, "service_name", bs.ServiceName, "Service identifier")
	flag.CommandLine.StringVar(&bs.LogLevel, "log_level", bs.LogLevel, "Logging Level")
	flag.CommandLine.BoolVar(&bs.LogToStdout, "stdout", bs.LogToStdout, "Log to stdout")
	flag.CommandLine.StringVar(&bs.MetricsDestination, "metrics", bs.MetricsDestination, "Metrics Configuration. none|syslog|stderr|graphite:127.0.0.1:2003")
	flag.CommandLine.StringVar(&bs.MetaListen, "metalisten", bs.MetaListen, "HTTP port to bind on")
	flag.CommandLine.StringVar(&bs.LegacyHealthCheckListen, "healthcheck", bs.LegacyHealthCheckListen, "Bind to this interface to allow healthchecks")

	flag.Parse()

	if *flag_version {
		fmt.Printf("%s: %s\nBuilt on %s %s (%s) \n", bs.ServiceName, bs.VersionInfo.Version,
			bs.VersionInfo.Platform, bs.VersionInfo.Distro, bs.VersionInfo.Date)
		os.Exit(0)
	}

	// validate our configuration
	if bs.ServiceName == "" {
		bs.Fail(fmt.Sprintf("Bad value for ServiceName [%s]", bs.ServiceName))
	}
}

// Perform some early initialization steps -- things it makes sense to do before callers start building/initializing
// anything from the actual service. Most things should probably be started from Run().  Called by Boilerplate().
func (bs *BaseServer) Init() {
	bs.InitLogger(bs.LogToStdout, bs.LogLevel)
	bs.InitMaxProcs()
	bs.InitMetrics()
}

func (bs *BaseServer) Fail(msg string) {
	if bs.Logger != nil {
		bs.Logger.Panic(bs.LogPrefix, msg)
	}
	fmt.Printf("%s\n", msg)
	os.Exit(FAILURE_CODE)
}

func (bs *BaseServer) WaitUntilReady(timeout time.Duration) {
	concurrent.WgWaitTimeout(&bs.waitGroup, timeout)
}

// Finish initializing and run until signaled otherwise. Spawns sub routines.
func (bs *BaseServer) Run(service Service) {
	bs.ctx, bs.cancel = context.WithCancel(context.WithValue(context.WithValue(context.Background(), readinessWaitGroupContextKey, &bs.waitGroup), subContextNameContextKey, "BaseServer.run"))

	bs.spawnPropsRefresh(bs.ctx)
	bs.spawnHealthCheck(bs.readyAwareSubContext(bs.ctx, "health check"), service)
	bs.spawnLegacyHealthCheck(bs.readyAwareSubContext(bs.ctx, "legacy health check"), service)
	bs.spawnMetaServer(bs.readyAwareSubContext(bs.ctx, "metaserver"), service)

	// run the actual service
	go func(ctx context.Context) {
		setReady(ctx)
		if err := service.Run(ctx); err != nil {
			bs.Fail(fmt.Sprintf("Cannot run service: %v", err))
		}
		bs.cancel()
	}(bs.readyAwareSubContext(bs.ctx, "service run goroutine"))

	s := make(chan os.Signal, 2)
	signal.Notify(s, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	setReady(bs.ctx) // goes with waitGroup.Add(1) in NewBaseServer
	for {
		select {
		case <-bs.ctx.Done():
			service.Close()
			bs.Logger.Info(bs.LogPrefix, "service.Close() called, now waiting for things to settle")
			time.Sleep(bs.ShutdownSettleTime) // Give everything enough time to settle.
			return
		case sig := <-s:
			switch sig {
			case syscall.SIGQUIT:
				bs.Shutdown("SIGQUIT")
			case syscall.SIGINT:
				bs.Shutdown("SIGINT")
			case syscall.SIGTERM:
				bs.Shutdown("SIGTERM")
			}
		}
	}
}

func (bs *BaseServer) Shutdown(reason string) {
	bs.Logger.Info(bs.LogPrefix, "Shutdown('%s')", reason)
	bs.cancel()
}

// Initialize logging.
func (bs *BaseServer) InitLogger(stdout bool, loglevel string) {

	bs.LogPrefix = bs.ServiceName + " "

	if stdout {
		logger.SetStdOut()
	}

	progSvcName := path.Base(os.Args[0])
	if progSvcName != bs.ServiceName {
		progSvcName = fmt.Sprintf("%s/%s", progSvcName, bs.ServiceName)
	}

	pid := os.Getpid()
	if pid > 10 {
		// Note about the above comparison: if our pid is super low, we're probably running inside docker and/or in a
		// context where pid is not likely to be very important, and we omit it.
		progSvcName = fmt.Sprintf("%s(%d)", progSvcName, pid)
	}

	if err := logger.SetLogName(fmt.Sprintf("%s ", progSvcName)); err != nil {
		bs.Fail("Cannot set log name for program")
	}
	if ll, ok := logger.CfgLevels[strings.ToLower(loglevel)]; !ok {
		bs.Fail("Unsupported log level: " + loglevel)
	} else {
		if bs.Logger = logger.New(ll); bs.Logger == nil {
			bs.Fail("Cannot start logger")
		}
	}
}

// Set the number of cpus this process can use.
func (bs *BaseServer) InitMaxProcs() {
	if nc, err := strconv.Atoi(os.Getenv(ENV_CH_NUM_CPU)); err == nil {
		runtime.GOMAXPROCS(nc)
		bs.Logger.Info(bs.LogPrefix, "Setting GOMAXPROCS to %d", nc)
	}
}

// Initialize metrics.
func (bs *BaseServer) InitMetrics() {
	tags := []string{
		"ver=" + bs.VersionInfo.Version,
		"svc=" + bs.ServiceName,
	}
	cmetrics.SetConf(bs.MetricsDestination, bs.Logger, bs.LogPrefix, bs.MetricsPrefix, nil, tags, nil, nil)
}

// Initialize our legacy health check.
func (bs *BaseServer) spawnHealthCheck(ctx context.Context, service Service) {
	bs.hce = NewHealthCheckExecutor(service, bs.HealthCheckStartupDelay, bs.HealthCheckPeriod, bs.HealthCheckTimeout)
	go bs.hce.Run(ctx)
}

// Start legacy healthcheck if needed. Called as part of Init()
func (bs *BaseServer) spawnLegacyHealthCheck(ctx context.Context, service Service) {
	setReady(ctx)
	if bs.LegacyHealthCheckListen != "" {
		// todo: check if the service implements healthcheck command handling and pass handleCmd accordingly
		go healthcheck.Run(bs.LegacyHealthCheckListen, service.GetStatus, nil, bs.Logger)
	}
}

func (bs *BaseServer) spawnMetaServer(ctx context.Context, service Service) {
	if bs.hce == nil {
		bs.Fail("initMetaServer: hce cannot be nil")
	}

	go func() {
		bs.metaServer = NewMetaServer(bs.MetaListen, bs.ServiceName, bs.VersionInfo, service, bs.Logger, bs.hce)
		if err := bs.metaServer.Run(ctx); err != nil {
			bs.Fail(fmt.Sprintf("Error running meta server: %+v", err))
		}
	}()
}

func (bs *BaseServer) spawnPropsRefresh(ctx context.Context) {
	go func() {
		for {
			select {
			case <-time.After(bs.PropsRefreshPeriod):
				bs.propertyService.Refresh()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (bs *BaseServer) readyAwareSubContext(ctx context.Context, name string) context.Context {
	val := ctx.Value(readinessWaitGroupContextKey)

	if val == nil {
		bs.Fail("Context is missing a value for readinessWaitGroupContextKey")
	}

	wg := val.(*sync.WaitGroup)
	wg.Add(1)

	// fmt.Printf("+ subcontext(%s) wg(%+v)\n", name, wg)
	return context.WithValue(context.WithValue(ctx, readinessWaitGroupContextKey, wg), subContextNameContextKey, name)
}

func setReady(ctx context.Context) {
	/* subCtxNameStr := "UNKNOWN"
	if val := ctx.Value(subContextNameContextKey); val != nil {
		subCtxNameStr = val.(string)
	}
	*/
	if val := ctx.Value(readinessWaitGroupContextKey); val != nil {
		wg := val.(*sync.WaitGroup)
		// fmt.Printf("- subcontext(%s) wg(%+v)\n", subCtxNameStr, wg)
		wg.Done()
	} /* else {
		 fmt.Printf("- subcontext(%s) wg(%+v)\n", subCtxNameStr, nil)
	} */
}
