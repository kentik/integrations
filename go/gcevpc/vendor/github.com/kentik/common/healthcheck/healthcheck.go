package healthcheck

import (
	"bufio"
	"fmt"
	"github.com/kentik/golog/logger"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime"
	"time"
)

const (
	LOG_PREFIX      = "[HealthCheck] "
	PPROF_BIND_ADDR = "PPROF_BIND_ADDR"
	INIT_TIME       = 10 * time.Second
)

var (
	GOOD = []byte("GOOD\n")
	BAD  = []byte("BAD\n")
)

func nilStatus() []byte {
	return GOOD
}

func nilCmd(cmd []byte) []byte {
	return []byte(fmt.Sprintf("Unknown command: %s\n", string(cmd)))
}

// GetMemStats returns a simple message describing the state of heap memory, as seen by the runtime.
// All values are in MB.
// - Sys:          bytes obtained from system
// - HeapSys:      bytes obtained from system
// - HeapAlloc:    bytes allocated and not yet freed
// - HeapIdle:     bytes in idle spans
// - HeapReleased: bytes released to the OS
func GetMemStats() string {
	mb := uint64(1000000)
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	return fmt.Sprintf("Sys: %d, HeapSys: %d, HeapAlloc: %d, HeapIdle: %d, HeapReleased: %d",
		memStats.Sys/mb, memStats.HeapSys/mb, memStats.HeapAlloc/mb, memStats.HeapIdle/mb, memStats.HeapReleased/mb)
}

func peekCmd(c net.Conn) ([]byte, error) {
	c.SetReadDeadline(time.Now().Add(time.Millisecond * 500))
	r := bufio.NewReader(c)

	if _, err := r.Peek(1); err == nil {
		// Extend read deadline
		c.SetReadDeadline(time.Now().Add(time.Second * 5))
		if buf, err := r.ReadBytes('\n'); err != nil {
			return nil, err
		} else {
			return buf[:len(buf)-1], nil
		}
	}
	return nil, nil
}

func Run(host string, statusReport func() []byte, handleCmd func([]byte) []byte, log *logger.Logger) {

	// Sleep for the time needed to settle things.
	time.Sleep(INIT_TIME)

	l, err := net.Listen("tcp", host)
	if err != nil {
		log.Error(LOG_PREFIX, "Error Binding to %s: %v", host, err)
		return
	}

	if statusReport == nil {
		statusReport = nilStatus
	}

	if handleCmd == nil {
		handleCmd = nilCmd
	}

	// Start up a pprof server as well.
	pprofBind := os.Getenv(PPROF_BIND_ADDR)
	if pprofBind != "" {
		go func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/debug/pprof/", pprof.Index)
			mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
			log.Info(LOG_PREFIX, "pprof listening at %s", pprofBind)
			log.Error(LOG_PREFIX, "%s", http.ListenAndServe(pprofBind, mux))
		}()
	}

	log.Info(LOG_PREFIX, "HC online at %s", host)
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Error(LOG_PREFIX, "Error Accepting request: %v", err)
			continue
		}

		if cmd, err := peekCmd(conn); err != nil {
			log.Error(LOG_PREFIX, "Error reading command: %s", err)
		} else if cmd != nil {
			resp := handleCmd(cmd)
			conn.Write(resp)
		} else {
			resp := statusReport()
			conn.Write(resp)
		}

		conn.Close()
	}

	log.Error(LOG_PREFIX, "HC Done at %s", host)
}
