package internal

import (
	"context"
	"log"

	blktk "github.com/Magicking/gethitihteg"
)

type key int

var blkKey key = 1
var ethRpcKey key = 2
var monitoringKey key = 3

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

func NewBLKToContext(ctx context.Context, wsURI, privateKey string) context.Context {
	blk, err := blktk.NewBlockchainContext(wsURI, privateKey, 5)
	if err != nil {
		log.Fatalf("Could not initialize blockchain context: %v", err)
	}
	return context.WithValue(ctx, blkKey, blk)
}

func BLKFromContext(ctx context.Context) (*blktk.BlockchainContext, bool) {
	blk, ok := ctx.Value(blkKey).(*blktk.BlockchainContext)
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