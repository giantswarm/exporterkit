package exporterkit

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

const (
	// DefaultAddress is the address the Exporter will run on if no address is configured.
	DefaultAddress = "0.0.0.0:8000"
)

// Config if the configuration to create an Exporter.
type Config struct {
	Address    string
	Collectors []prometheus.Collector
	Logger     micrologger.Logger
}

// Exporter runs a slice of Prometheus Collectors.
type Exporter struct {
	address    string
	collectors []prometheus.Collector
	logger     micrologger.Logger
}

// New creates a new Exporter, given a Config.
func New(config Config) (*Exporter, error) {
	if config.Collectors == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Collectors must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.Address == "" {
		config.Address = DefaultAddress
	}

	exporter := Exporter{
		address:    config.Address,
		collectors: config.Collectors,
		logger:     config.Logger,
	}

	return &exporter, nil
}

// Run starts the Exporter.
func (e *Exporter) Run() {
	e.logger.Log("level", "info", "message", fmt.Sprintf("starting exporter on %s", e.address))

	prometheus.MustRegister(e.collectors...)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok\n")
	}))

	http.ListenAndServe(e.address, nil)
}
