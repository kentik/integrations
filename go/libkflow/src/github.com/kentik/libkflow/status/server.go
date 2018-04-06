package status

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/kentik/libkflow/log"
	"github.com/kentik/libkflow/metrics"
)

type Server struct {
	metrics *metrics.Metrics
	router  *http.ServeMux
	server  *http.Server
}

func NewServer(host string, port int) *Server {
	router := &http.ServeMux{}
	server := &http.Server{
		Addr:         net.JoinHostPort(host, strconv.Itoa(port)),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	return &Server{
		router: router,
		server: server,
	}
}

func (s *Server) Start(m *metrics.Metrics) {
	log.Debugf("status server at %s", s.server.Addr)
	s.metrics = m
	s.router.Handle("/v1/status", s)
	err := s.server.ListenAndServe()
	log.Debugf("status server error: %s", err)
}

type Status struct {
	FlowsIn  Stats `json:"flows-in"`
	FlowsOut Stats `json:"flows-out"`
}

type Stats struct {
	Count int64   `json:"count"`
	Rate1 float64 `json:"1m.rate"`
	Rate5 float64 `json:"5m.rate"`
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := Status{
		FlowsIn: Stats{
			Count: s.metrics.TotalFlowsIn.Count(),
			Rate1: s.metrics.TotalFlowsIn.Rate1(),
			Rate5: s.metrics.TotalFlowsIn.Rate5(),
		},
		FlowsOut: Stats{
			Count: s.metrics.TotalFlowsOut.Count(),
			Rate1: s.metrics.TotalFlowsOut.Rate1(),
			Rate5: s.metrics.TotalFlowsOut.Rate5(),
		},
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(&status)

	w.Header().Set("Content-Type", "application/json")
}
