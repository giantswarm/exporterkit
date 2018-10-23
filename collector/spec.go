package collector

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
)

// Interface defines how a collector implementation should look like.
type Interface interface {
	// Boot is used to initial and register the collector. Boot must be allowed to
	// be called multiple times and thus must be idempotent. Usual implementations
	// could make use of sync.Once. See also
	// https://godoc.org/github.com/prometheus/client_golang/prometheus#Register.
	Boot(ctx context.Context) error
	// Collect should align with the monitoring system's implementation
	// requirements. In this case Prometheus. See also
	// https://godoc.org/github.com/prometheus/client_golang/prometheus#Collector.
	// The difference here is that this specific interface provides additional
	// error handling capabilities.
	CollectWithError(ch chan<- prometheus.Metric) error
	// Describe should align with the monitoring system's implementation
	// requirements. In this case Prometheus. See also
	// https://godoc.org/github.com/prometheus/client_golang/prometheus#Collector.
	// The difference here is that this specific interface provides additional
	// error handling capabilities.
	DescribeWithError(ch chan<- *prometheus.Desc) error
}
