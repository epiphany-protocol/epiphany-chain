package blockchain

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// Metrics represents the blockchain metrics
type Metrics struct {
	//TPS of recent minute
	TPSRecentMinute metrics.Gauge
	//TPS of recent hour
	TPSRecentHour metrics.Gauge
}

// GetPrometheusMetrics return the blockchain metrics instance
func GetPrometheusMetrics(namespace string, labelsWithValues ...string) *Metrics {
	labels := []string{}

	for i := 0; i < len(labelsWithValues); i += 2 {
		labels = append(labels, labelsWithValues[i])
	}

	return &Metrics{
		TPSRecentMinute: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "tps_recent_minute",
			Help:      "TPS of recent minute.",
		}, labels).With(labelsWithValues...),
		TPSRecentHour: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "tps_recent_hour",
			Help:      "TPS of recent hour.",
		}, labels).With(labelsWithValues...),
	}
}

// NilMetrics will return the non operational metrics
func NilMetrics() *Metrics {
	return &Metrics{
		TPSRecentMinute: discard.NewGauge(),
		TPSRecentHour:   discard.NewGauge(),
	}
}
