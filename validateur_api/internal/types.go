package internal

import (
	"context"
	"log"

	blktk "github.com/Magicking/gethitihteg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type key int

var blkKey key = 1
var ethRpcKey key = 2
var monitoringKey key = 3

type NodeConnector struct {
	client  *ethclient.Client
	rawurl  string
	try_max int
}

// Connector to get Tx
type BlockchainContext struct {
	NC *NodeConnector
}

func (nc *NodeConnector) GetTransaction(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	tx, isPending, err = nc.client.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, false, err
	}
	return tx, isPending, nil
}

func NewCCToContext(ctx context.Context, wsURI string) context.Context {
	cc, err := blktk.NewClientConnector(wsURI, 3)
	if err != nil {
		log.Fatalf("Could not initialize client context: %v", err)
	}
	return context.WithValue(ctx, ethRpcKey, cc)
}

func CCFromContext(ctx context.Context) (*blktk.ClientConnector, bool) {
	cc, ok := ctx.Value(ethRpcKey).(*blktk.ClientConnector)
	return cc, ok
}

func NewBLKToContext(ctx context.Context, wsURI string) context.Context {
	client, err := ethclient.Dial(wsURI)
	if err != nil {
		log.Fatalf("Could not initialize blockchain context: %v", err)
	}

	nodeConnector := NodeConnector{client: client, rawurl: wsURI}

	blk := &BlockchainContext{NC: &nodeConnector}

	return context.WithValue(ctx, blkKey, blk)
}

func BLKFromContext(ctx context.Context) (*BlockchainContext, bool) {
	blk, ok := ctx.Value(blkKey).(*BlockchainContext)

	return blk, ok
}

func NewMonitoringToContext(ctx context.Context, nodeAddress string, lockedAddress string) context.Context {
	mn := InitMonitoring(nodeAddress, lockedAddress)
	if (MonitoringEnv{}) == mn {
		log.Fatalf("Could not initialize monitoring cont: %v", mn)
	}
	return context.WithValue(ctx, monitoringKey, mn)
}

func MonitoringFromContext(ctx context.Context) (MonitoringEnv, bool) {
	mn, ok := ctx.Value(monitoringKey).(MonitoringEnv)
	return mn, ok
}