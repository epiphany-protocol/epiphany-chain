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
	//Average block execution period in recent 5 minutes
	AvrgBlockPeriodRecent5Min metrics.Gauge
	//Average block execution period in recent hour
	AvrgBlockPeriodRecentHour metrics.Gauge
	//Average transaction execution period in recent 5 minutes
	AvrgTxPeriodRecent5Min metrics.Gauge
	//Average transaction execution period in recent hour
	AvrgTxPeriodRecentHour metrics.Gauge

	//Height of blockchain
	BlockHeight metrics.Gauge
	//Last block execution period
	BlockPeriod metrics.Gauge
	//TPS of database
	TPSDB metrics.Gauge
	//Database time usage
	DBPeriod metrics.Gauge

	//Occured forks
	Forks metrics.Counter

	//Error Messages occured
	ErrorMessages metrics.Counter
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
		AvrgBlockPeriodRecent5Min: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "avrg_block_period_recent_5_min",
			Help:      "Average block execution period in recent 5 minutes.",
		}, labels).With(labelsWithValues...),
		AvrgBlockPeriodRecentHour: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "avrg_block_period_recent_hour",
			Help:      "Average block execution period in recent hour.",
		}, labels).With(labelsWithValues...),
		AvrgTxPeriodRecent5Min: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "avrg_tx_period_recent_5_min",
			Help:      "Average transaction execution period in recent 5 minutes.",
		}, labels).With(labelsWithValues...),
		AvrgTxPeriodRecentHour: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "avrg_tx_period_recent_hour",
			Help:      "Average transaction execution period in recent hour.",
		}, labels).With(labelsWithValues...),
		BlockHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "block_height",
			Help:      "Height of blockchain.",
		}, labels).With(labelsWithValues...),
		BlockPeriod: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "block_period",
			Help:      "Last block execution period.",
		}, labels).With(labelsWithValues...),
		TPSDB: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "tps_db",
			Help:      "TPS of database.",
		}, labels).With(labelsWithValues...),
		DBPeriod: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "db_period",
			Help:      "Time usage of DB persisting blocks.",
		}, labels).With(labelsWithValues...),
		Forks: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "forks",
			Help:      "Occured forks.",
		}, labels).With(labelsWithValues...),
		ErrorMessages: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "blockchain",
			Name:      "error_messages",
			Help:      "Error Messages occured.",
		}, labels).With(labelsWithValues...),
	}
}

// NilMetrics will return the non operational metrics
func NilMetrics() *Metrics {
	return &Metrics{
		TPSRecentMinute:           discard.NewGauge(),
		TPSRecentHour:             discard.NewGauge(),
		AvrgBlockPeriodRecent5Min: discard.NewGauge(),
		AvrgBlockPeriodRecentHour: discard.NewGauge(),
		AvrgTxPeriodRecent5Min:    discard.NewGauge(),
		AvrgTxPeriodRecentHour:    discard.NewGauge(),
		BlockHeight:               discard.NewGauge(),
		BlockPeriod:               discard.NewGauge(),
		TPSDB:                     discard.NewGauge(),
		DBPeriod:                  discard.NewGauge(),
		Forks:                     discard.NewCounter(),
		ErrorMessages:             discard.NewCounter(),
	}
}
