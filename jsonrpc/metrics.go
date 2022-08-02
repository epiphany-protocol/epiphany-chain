package jsonrpc

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// Metrics represents the consensus metrics
type Metrics struct {
	//JSONRPC calls
	JsonRPCCalls metrics.Counter

	//JSONRPC call time period
	JsonRPCCallTime metrics.Gauge

	//Error Messages occured
	ErrorMessages metrics.Counter
}

// GetPrometheusMetrics return the consensus metrics instance
func GetPrometheusMetrics(namespace string, labelsWithValues ...string) *Metrics {
	labels := []string{}

	for i := 0; i < len(labelsWithValues); i += 2 {
		labels = append(labels, labelsWithValues[i])
	}

	return &Metrics{
		JsonRPCCalls: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "jsonrpc",
			Name:      "jsonrpc_calls",
			Help:      "JSONRPC calls.",
		}, labels).With(labelsWithValues...),
		JsonRPCCallTime: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "jsonrpc",
			Name:      "jsonrpc_call_time",
			Help:      "JSONRPC call time period",
		}, labels).With(labelsWithValues...),
		ErrorMessages: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "jsonrpc",
			Name:      "error_messages",
			Help:      "Error Messages occured.",
		}, labels).With(labelsWithValues...),
	}
}

// NilMetrics will return the non operational metrics
func NilMetrics() *Metrics {
	return &Metrics{
		JsonRPCCalls:    discard.NewCounter(),
		JsonRPCCallTime: discard.NewGauge(),
		ErrorMessages:   discard.NewCounter(),
	}
}
