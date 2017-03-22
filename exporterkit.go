package exporterkit

import (
	"net/http"
	"os"

	kithttp "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"

	"github.com/giantswarm/microkit/logger"
	"github.com/giantswarm/microkit/server"
)

// Handler represents a function that is called by an exporter.
type Handler func() error

// Config represents the configuration for an exporter.
type Config struct {
	Name        string
	Description string

	Handler func(logger.Logger) error
}

// Exporter represents a prometheus exporter.
type Exporter interface {
	// Run starts the exporter.
	Run()
}

// NewExporter creates an Exporter, given a Handler.
func NewExporter(config Config) Exporter {
	return &exporter{
		config: config,
	}
}

// exporter is the basic implementation of Exporter.
type exporter struct {
	config Config
}

func errorEncoder() kithttp.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Run starts the exporter.
func (e *exporter) Run() {
	loggerConfig := logger.DefaultConfig()
	loggerConfig.IOWriter = os.Stdout

	newLogger, err := logger.New(loggerConfig)
	if err != nil {
		panic(err)
	}

	serverConfig := server.DefaultConfig()
	serverConfig.Logger = newLogger
	serverConfig.Endpoints = []server.Endpoint{}

	server, err := server.New(serverConfig)
	if err != nil {
		panic(err)
	}

	server.Boot()
}
