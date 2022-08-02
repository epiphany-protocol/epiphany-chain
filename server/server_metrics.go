package server

import (
	"github.com/0xPolygon/polygon-edge/blockchain"
	"github.com/0xPolygon/polygon-edge/consensus"
	"github.com/0xPolygon/polygon-edge/jsonrpc"
	"github.com/0xPolygon/polygon-edge/network"
	"github.com/0xPolygon/polygon-edge/network/discovery"
	"github.com/0xPolygon/polygon-edge/protocol"
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
	jsonrpc    *jsonrpc.Metrics
	protocol   *protocol.Metrics
	discovery  *discovery.Metrics
	server     *Metrics
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
			jsonrpc:    jsonrpc.GetPrometheusMetrics(nameSpace, "chain_id", chainID),
			protocol:   protocol.GetPrometheusMetrics(nameSpace, "chain_id", chainID),
			discovery:  discovery.GetPrometheusMetrics(nameSpace, "chain_id", chainID),
			server:     GetPrometheusMetrics(nameSpace, "chain_id", chainID),
		}
	}

	return &serverMetrics{
		consensus:  consensus.NilMetrics(),
		network:    network.NilMetrics(),
		txpool:     txpool.NilMetrics(),
		blockchain: blockchain.NilMetrics(),
		state:      state.NilMetrics(),
		jsonrpc:    jsonrpc.NilMetrics(),
		protocol:   protocol.NilMetrics(),
		discovery:  discovery.NilMetrics(),
		server:     NilMetrics(),
	}
}
