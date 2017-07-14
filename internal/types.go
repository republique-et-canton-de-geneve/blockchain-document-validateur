package internal

import (
	"context"
	"github.com/jinzhu/gorm"
	"log"
)

type key int

var dbKey key = 0

//var blkKey key = 1

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
