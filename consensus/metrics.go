package consensus

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// Metrics represents the consensus metrics
type Metrics struct {
	// No.of validators
	Validators metrics.Gauge
	// No.of rounds
	Rounds metrics.Gauge
	// No.of transactions in the block
	NumTxs metrics.Gauge

	//Time between current block and the previous block in seconds
	BlockInterval metrics.Gauge

	//Blocks latency below 500ms
	BlockLatencyBelow500ms metrics.Counter

	//Blocks latency below 1s
	BlockLatencyBelow1s metrics.Counter

	//Blocks latency below 2s
	BlockLatencyBelow2s metrics.Counter

	//Blocks latency below 3s
	BlockLatencyBelow3s metrics.Counter

	//Round changes
	RoundChanges metrics.Counter
}

// GetPrometheusMetrics return the consensus metrics instance
func GetPrometheusMetrics(namespace string, labelsWithValues ...string) *Metrics {
	labels := []string{}

	for i := 0; i < len(labelsWithValues); i += 2 {
		labels = append(labels, labelsWithValues[i])
	}

	return &Metrics{
		Validators: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "consensus",
			Name:      "validators",
			Help:      "Number of validators.",
		}, labels).With(labelsWithValues...),
		Rounds: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "consensus",
			Name:      "rounds",
			Help:      "Number of rounds.",
		}, labels).With(labelsWithValues...),
		NumTxs: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "consensus",
			Name:      "num_txs",
			Help:      "Number of transactions.",
		}, labels).With(labelsWithValues...),

		BlockInterval: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "consensus",
			Name:      "block_interval",
			Help:      "Time between current block and the previous block in seconds.",
		}, labels).With(labelsWithValues...),

		BlockLatencyBelow500ms: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "consensus",
			Name:      "block_latency_below_500ms",
			Help:      "Blocks latency below 500ms.",
		}, labels).With(labelsWithValues...),

		BlockLatencyBelow1s: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "consensus",
			Name:      "block_latency_below_1s",
			Help:      "Blocks latency below 1s.",
		}, labels).With(labelsWithValues...),

		BlockLatencyBelow2s: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "consensus",
			Name:      "block_latency_below_2s",
			Help:      "Blocks latency below 2s.",
		}, labels).With(labelsWithValues...),

		BlockLatencyBelow3s: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "consensus",
			Name:      "block_latency_below_3s",
			Help:      "Blocks latency below 3s.",
		}, labels).With(labelsWithValues...),

		RoundChanges: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "consensus",
			Name:      "round_changes",
			Help:      "Round Changes.",
		}, labels).With(labelsWithValues...),
	}
}

// NilMetrics will return the non operational metrics
func NilMetrics() *Metrics {
	return &Metrics{
		Validators:             discard.NewGauge(),
		Rounds:                 discard.NewGauge(),
		NumTxs:                 discard.NewGauge(),
		BlockInterval:          discard.NewGauge(),
		BlockLatencyBelow500ms: discard.NewCounter(),
		BlockLatencyBelow1s:    discard.NewCounter(),
		BlockLatencyBelow2s:    discard.NewCounter(),
		BlockLatencyBelow3s:    discard.NewCounter(),
		RoundChanges:           discard.NewCounter(),
	}
}
