package internal

import (
	"context"
	blktk "github.com/Magicking/gethitihteg"
	"github.com/jinzhu/gorm"
	"log"
)

type key int

var dbKey key = 0
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

func NewDBToContext(ctx context.Context, dbDsn string) context.Context {
	db, err := InitDatabase(dbDsn)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	return context.WithValue(ctx, dbKey, db)
}

func DBFromContext(ctx context.Context) (*gorm.DB, bool) {
	db, ok := ctx.Value(dbKey).(*gorm.DB)
	return db, ok
}
