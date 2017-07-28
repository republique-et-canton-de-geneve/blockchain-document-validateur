package internal

import (
	"context"
	blktk "github.com/Magicking/gethitihteg"
	"log"
)

type key int

var blkKey key = 1
var ethRpcKey key = 2

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
	blk, err := blktk.NewBlockchainContext(wsURI, privateKey)
	if err != nil {
		log.Fatalf("Could not initialize blockchain context: %v", err)
	}
	return context.WithValue(ctx, blkKey, blk)
}

func BLKFromContext(ctx context.Context) (*blktk.BlockchainContext, bool) {
	blk, ok := ctx.Value(blkKey).(*blktk.BlockchainContext)
	return blk, ok
}
