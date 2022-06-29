package server

import (
	"github.com/0xPolygon/polygon-edge/blockchain"
	"github.com/0xPolygon/polygon-edge/consensus"
	"github.com/0xPolygon/polygon-edge/network"
	"github.com/0xPolygon/polygon-edge/state"
	"github.com/0xPolygon/polygon-edge/txpool"
)

// serverMetrics holds the metric instances of all sub systems
type serverMetrics struct {
	consensus  *consensus.Metrics
	network    *network.Metrics
	txpool     *txpool.Metrics
	blockchain *blockchain.Metrics
	state      *state.Metrics
}

// metricProvider serverMetric instance for the given ChainID and nameSpace
func metricProvider(nameSpace string, chainID string, metricsRequired bool) *serverMetrics {
	if metricsRequired {
		return &serverMetrics{
			consensus:  consensus.GetPrometheusMetrics(nameSpace, "chain_id", chainID),
			network:    network.GetPrometheusMetrics(nameSpace, "chain_id", chainID),
			txpool:     txpool.GetPrometheusMetrics(nameSpace, "chain_id", chainID),
			blockchain: blockchain.GetPrometheusMetrics(nameSpace, "chain_id", chainID),
			state:      state.GetPrometheusMetrics(nameSpace, "chain_id", chainID),
		}
	}

	return &serverMetrics{
		consensus:  consensus.NilMetrics(),
		network:    network.NilMetrics(),
		txpool:     txpool.NilMetrics(),
		blockchain: blockchain.NilMetrics(),
		state:      state.NilMetrics(),
	}
}
