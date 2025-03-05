package exporterkit

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/giantswarm/microendpoint/endpoint/healthz"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// DefaultAddress is the address the Exporter will run on if no address is configured.
	DefaultAddress = "http://0.0.0.0:8000"
)

// Config if the configuration to create an Exporter.
type Config struct {
	Address        string
	Collectors     []prometheus.Collector
	ExtraEndpoints []server.Endpoint
	Logger         micrologger.Logger
}

// Exporter runs a slice of Prometheus Collectors.
type Exporter struct {
	address        string
	collectors     []prometheus.Collector
	extraEndpoints []server.Endpoint
	logger         micrologger.Logger
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

	for _, e := range config.ExtraEndpoints {
		if e.Path() == healthz.Path {
			return nil, microerror.Maskf(invalidConfigError, "%T.ExtraEndpoints: endpoints with path %q can not be added", config, healthz.Path)
		}
	}

	exporter := Exporter{
		address:        config.Address,
		collectors:     config.Collectors,
		extraEndpoints: config.ExtraEndpoints,
		logger:         config.Logger,
	}

	return &exporter, nil
}

// Run starts the Exporter.
func (e *Exporter) Run() {
	var err error

	var healthzEndpoint *healthz.Endpoint
	{
		c := healthz.Config{
			Logger: e.logger,
		}
		healthzEndpoint, err = healthz.New(c)
		if err != nil {
			panic(fmt.Sprintf("%#v\n", err))
		}
	}

	var newServer server.Server
	{
		c := server.Config{
			EnableDebugServer: true,
			Endpoints:         append(e.extraEndpoints, healthzEndpoint),
			ListenAddress:     e.address,
			Logger:            e.logger,
		}

		newServer, err = server.New(c)
		if err != nil {
			panic(fmt.Sprintf("%#v\n", err))
		}
	}

	prometheus.MustRegister(e.collectors...)

	go newServer.Boot()

	listener := make(chan os.Signal, 2)
	signal.Notify(listener, os.Interrupt, syscall.SIGTERM)

	<-listener

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			newServer.Shutdown()
		}()

		os.Exit(0)
	}()

	<-listener

	os.Exit(0)
}
