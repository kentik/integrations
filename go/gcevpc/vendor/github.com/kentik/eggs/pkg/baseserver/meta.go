package baseserver

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime"
	"strconv"
	"strings"

	"context"
	"fmt"
	"github.com/kentik/eggs/pkg/version"

	go_metrics "github.com/kentik/go-metrics"
	"github.com/kentik/golog/logger"
	gopsutil_cpu "github.com/shirou/gopsutil/cpu"
	gopsutil_disk "github.com/shirou/gopsutil/disk"
	gopsutil_host "github.com/shirou/gopsutil/host"
	gopsutil_load "github.com/shirou/gopsutil/load"
	gopsutil_mem "github.com/shirou/gopsutil/mem"
	gopsutil_net "github.com/shirou/gopsutil/net"
	gopsutil_process "github.com/shirou/gopsutil/process"
)

const (
	MS_LOG_PREFIX = "baseserver.metaserver "
)

// All lowercase here please
var ENV_BLACKLIST = []string{
	"password",
	"token",
	"credential",
	"pg_connection",
	"pg_write_connection",
	"mailer",
	"alert_connection",
}

type MetaServer struct {
	listen      string
	mux         *http.ServeMux
	service     Service
	serviceName string
	version     version.VersionInfo
	log         *logger.Logger
	hce         *HealthCheckExecutor
	listenAddr  net.Addr
}

func NewMetaServer(listen string, serviceName string, version version.VersionInfo, service Service, log *logger.Logger, hce *HealthCheckExecutor) *MetaServer {
	ms := &MetaServer{
		listen:      listen,
		mux:         http.NewServeMux(),
		serviceName: serviceName,
		version:     version,
		service:     service,
		log:         log,
		hce:         hce,
	}

	// status, version, env etc
	ms.mux.HandleFunc("/sys", ms.endpoint_system)
	ms.mux.HandleFunc("/ps/", ms.endpoint_ps)
	ms.mux.HandleFunc("/env", ms.endpoint_environment)
	ms.mux.HandleFunc("/version", ms.endpoint_version)

	// metrics
	ms.mux.HandleFunc("/metrics", ms.endpoint_metrics)

	// healthcheck
	ms.mux.HandleFunc("/healthcheck", ms.endpoint_healthcheck)
	ms.mux.HandleFunc("/hc", ms.endpoint_healthcheck)

	// properties
	ms.mux.HandleFunc("/prop/", ms.endpoint_props)

	// debug/memstats
	ms.mux.HandleFunc("/debug/memstats", ms.endpoint_memstats)

	// debug/pprof
	ms.mux.HandleFunc("/debug/pprof/", pprof.Index)
	ms.mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	ms.mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	ms.mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	ms.mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// info
	ms.mux.HandleFunc("/service/info", service.HttpInfo)

	return ms
}

func (ms *MetaServer) Run(ctx context.Context) error {
	// we do this instead of just calling http.ListenAndServe so we can get the listening address, in case that
	// was not configured and we're using a random one. Drawback is that we lose keepalives.
	server := &http.Server{Addr: ms.listen, Handler: ms.mux}
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return err
	}
	ms.listenAddr = ln.Addr()
	ms.log.Infof(MS_LOG_PREFIX, "Listening on %v", ms.listenAddr) // nolint
	setReady(ctx)
	return server.Serve(ln)
}

func (ms *MetaServer) writeCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Server", "kentik-baseserver-metaserver")
	w.Header().Set("X-Kentik-Service", ms.serviceName)
	w.Header().Set("X-Kentik-Version", ms.version.Version)
	w.Header().Set("Content-Type", "application/json")

}

func (ms *MetaServer) writeJson(r *http.Request, w http.ResponseWriter, payload interface{}) {
	var (
		jsonResult []byte
		err        error
	)

	if r.URL.Query().Get("indent") != "0" {
		jsonResult, err = json.MarshalIndent(payload, "", " ")
	} else {
		jsonResult, err = json.Marshal(payload)
	}

	if err != nil {
		ms.log.Errorf(MS_LOG_PREFIX, "Could not serialize payload: %v", err) // nolint
		ms.writeError(w, err, http.StatusInternalServerError)
	} else {
		_, err = w.Write(jsonResult)
		if err != nil {
			ms.log.Errorf(MS_LOG_PREFIX, "Could not write payload: %v", err) // nolint
		}
	}
}

func (ms *MetaServer) writePlain(r *http.Request, w http.ResponseWriter, payload interface{}) {
	if _, err := w.Write([]byte(fmt.Sprintf("%+v\n", payload))); err != nil {
		ms.log.Errorf(MS_LOG_PREFIX, "Could not write payload: %v", err) // nolint
	}
}

func (ms *MetaServer) writeResponse(r *http.Request, w http.ResponseWriter, payload interface{}) {
	if r.URL.Query().Get("plain") != "" {
		ms.writePlain(r, w, payload)
	} else {
		ms.writeJson(r, w, payload)
	}
}

func (ms *MetaServer) writeError(w http.ResponseWriter, errToWrite error, code int) {
	w.WriteHeader(code)
	errorPayload := map[string]interface{}{
		"error": fmt.Sprintf("%+v", errToWrite),
	}
	var bytesToWrite []byte
	if jsonBytes, err := json.Marshal(errorPayload); err != nil {
		bytesToWrite = []byte("{}")
	} else {
		bytesToWrite = jsonBytes
	}

	if _, err := w.Write(bytesToWrite); err != nil {
		ms.log.Errorf(MS_LOG_PREFIX, "Could not write payload: %v", err) // nolint
	}
}

func getRedactedEnvironment() (map[string]string, []string) {
	env := os.Environ()
	envmap := make(map[string]string)
	redacted := make([]string, 0, len(env))

	for _, line := range env {
		parts := strings.SplitN(line, "=", 2)
		name := parts[0]
		value := parts[1]

		lowerCaseValue := strings.ToLower(value)
		redact := false
		for _, blacklisted := range ENV_BLACKLIST {
			if strings.Contains(lowerCaseValue, blacklisted) {
				redact = true
				break
			}
		}
		if redact {
			redacted = append(redacted, name)
		} else {
			envmap[name] = value
		}
	}
	return envmap, redacted
}

func (ms *MetaServer) endpoint_environment(w http.ResponseWriter, r *http.Request) {
	ms.writeCommonHeaders(w)

	envmap, redacted := getRedactedEnvironment()
	result := map[string]interface{}{
		"env":      envmap,
		"redacted": redacted,
	}

	ms.writeResponse(r, w, result)
}

func (ms *MetaServer) endpoint_system(w http.ResponseWriter, r *http.Request) {
	ms.writeCommonHeaders(w)
	cpuInfo, _ := gopsutil_cpu.Info()
	diskInfo, _ := gopsutil_disk.Partitions(true)
	hostInfo, _ := gopsutil_host.Info()
	loadInfo, _ := gopsutil_load.Avg()
	memInfo, _ := gopsutil_mem.VirtualMemory()
	netInfo, _ := gopsutil_net.IOCounters(true)
	result := map[string]interface{}{
		"cpu":  cpuInfo,
		"disk": diskInfo,
		"host": hostInfo,
		"load": loadInfo,
		"mem":  memInfo,
		"net":  netInfo,
	}
	ms.writeResponse(r, w, result)
}

func (ms *MetaServer) endpoint_ps(w http.ResponseWriter, r *http.Request) {
	ms.writeCommonHeaders(w)

	pid, err := strconv.Atoi(r.URL.Path[len("/ps/"):])
	if err != nil {
		pid = os.Getpid()
	}

	var processInfo = make(map[string]interface{})
	if process, err := gopsutil_process.NewProcess(int32(pid)); err == nil {
		processInfo["name"], _ = process.Name()
		processInfo["exe"], _ = process.Exe()
		processInfo["cmdline"], _ = process.Cmdline()
		processInfo["cpuAffinity"], _ = process.CPUAffinity()
		processInfo["createTime"], _ = process.CreateTime()
		processInfo["numFDs"], _ = process.NumFDs()
		processInfo["numThreads"], _ = process.NumThreads()
		processInfo["memInfo"], _ = process.MemoryInfo()
		processInfo["memInfoEx"], _ = process.MemoryInfoEx()
		processInfo["netIOCounters"], _ = process.NetIOCounters(true)
		processInfo["mmap"], _ = process.MemoryMaps(true)
	}
	ms.writeResponse(r, w, processInfo)
}

func (ms *MetaServer) endpoint_version(w http.ResponseWriter, r *http.Request) {
	ms.writeCommonHeaders(w)
	ms.writeResponse(r, w, ms.version)
}

func (ms *MetaServer) endpoint_props(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ms.writeCommonHeaders(w)
		propName := r.URL.Path[len("/prop/"):]
		propValue := GetGlobalBaseServer().GetPropertyService().GetString(propName, "")
		if propValue.FromFallback() {
			ms.writeError(w, fmt.Errorf("Property '%s' is not defined", propName), http.StatusNotFound)
		} else {
			ms.writeResponse(r, w, propValue.String())
		}
	case "PATCH":
		ms.writeCommonHeaders(w)
		GetGlobalBaseServer().GetPropertyService().Refresh()
	}

}

func (ms *MetaServer) endpoint_memstats(w http.ResponseWriter, r *http.Request) {
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	ms.writeCommonHeaders(w)
	ms.writeResponse(r, w, memStats)
}

func (ms *MetaServer) endpoint_healthcheck(w http.ResponseWriter, r *http.Request) {
	result := ms.hce.GetResult()
	ms.writeCommonHeaders(w)
	if result.Success {
		ms.log.Debug(MS_LOG_PREFIX, "Healthcheck success") // nolint
		w.WriteHeader(http.StatusOK)
	} else {
		ms.log.Infof(MS_LOG_PREFIX, "Healthcheck failed (%v), returning 500", result) // nolint
		w.WriteHeader(http.StatusInternalServerError)
	}
	ms.writeResponse(r, w, result)
}

func _map(elems ...interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i, elem := range elems {
		if i%2 == 0 {
			key := elem.(string)
			value := elems[i+1]
			result[key] = value
		}
	}
	return result
}

func (ms *MetaServer) endpoint_metrics(w http.ResponseWriter, r *http.Request) {
	filter := r.FormValue("f")

	result := make(map[string]interface{})
	go_metrics.Each(func(name string, i interface{}) {

		if !strings.Contains(name, filter) {
			return
		}

		switch metric := i.(type) {

		case go_metrics.Healthcheck:
			ms.log.Error(MS_LOG_PREFIX, "Not expecting to see healchchecks here: %v", metric) // nolint
			return

		case go_metrics.Counter:
			result[name] = _map("type", "counter", "count", metric.Count())

		case go_metrics.Gauge:
			result[name] = _map("type", "gauge", "count", metric.Value())

		case go_metrics.GaugeFloat64:
			result[name] = _map("type", "gauge", "count", metric.Value())

		case go_metrics.Histogram:
			h := metric.Snapshot()
			ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			result[name] = _map("type", "histogram",
				"count", h.Count(),
				"min", h.Min(),
				"max", h.Max(),
				"mean", h.Mean(),
				"stddev", h.StdDev(),
				"p50", ps[0],
				"p75", ps[1],
				"p95", ps[2],
				"p99", ps[3],
				"p999", ps[4],
			)

		case go_metrics.Meter:
			m := metric.Snapshot()
			result[name] = _map("type", "meter",
				"count", m.Count(),
				"rate1", m.Rate1(),
				"rate5", m.Rate5(),
				"rate15", m.Rate15(),
				"mean", m.RateMean(),
			)

		case go_metrics.Timer:
			t := metric.Snapshot()
			ps := t.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			result[name] = _map("type", "timer",
				"count", t.Count(),
				"min", t.Min(),
				"max", t.Max(),
				"mean", t.Mean(),
				"stddev", t.StdDev(),
				"p50", ps[0],
				"p75", ps[1],
				"p95", ps[2],
				"p99", ps[3],
				"p999", ps[4],
				"rate1", t.Rate1(),
				"rate5", t.Rate5(),
				"rate15", t.Rate15(),
				"mean", t.RateMean(),
			)
		}
	})
	ms.writeCommonHeaders(w)
	ms.writeResponse(r, w, result)
}
