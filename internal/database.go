package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Receipt struct {
	gorm.Model
	TargetHash      string
	TransactionHash string
	Filename        string
	JSONData        []byte
}

func InsertReceipt(ctx context.Context, filename string, rcpt *Chainpoint) error {
	jsonData, err := json.Marshal(rcpt)
	if err != nil {
		return err
	}

	var tx_hash string

	if len(rcpt.Anchors) > 0 {
		tx_hash = rcpt.Anchors[0].SourceID
	}
	_rcpt := Receipt{
		TargetHash:      rcpt.TargetHash,
		JSONData:        jsonData,
		TransactionHash: tx_hash,
		Filename:        filename,
	}

	db, ok := DBFromContext(ctx)
	if !ok {
		return fmt.Errorf("Could not obtain DB from Context")
	}
	if err := db.Create(&_rcpt).Error; err != nil {
		return err
	}

	log.Printf("DEBUG: InsertReceipt: %v", _rcpt)

	return nil
}

func GetReceiptByHash(ctx context.Context, hash string) (*Receipt, bool, error) {
	var rcpt Receipt
	db, ok := DBFromContext(ctx)
	if !ok {
		return nil, false, fmt.Errorf("Could not obtain DB from Context")
	}
	cursor := db.Where(Receipt{TargetHash: hash})
	if cursor.Error != nil {
		return nil, false, fmt.Errorf("Error for TargetHash (%v): %v", hash, cursor.Error)
	}
	if cursor.Last(&rcpt).RecordNotFound() {
		return nil, false, nil
	}
	return &rcpt, true, nil
}

func GetAllReceipts(ctx context.Context) ([]Receipt, error) {
	var receipts []Receipt

	db, ok := DBFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Could not obtain DB from Context")
	}
	if db.Find(&receipts).RecordNotFound() {
		return nil, fmt.Errorf("RecordNotFound: No receipts in database")
	}

	return receipts, nil
}

func InitDatabase(dbDsn string) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	for i := 1; i < 10; i++ {
		db, err = gorm.Open("postgres", dbDsn)
		if err == nil || i == 10 {
			break
		}
		sleep := (2 << uint(i)) * time.Second
		log.Printf("Could not connect to DB: %v", err)
		log.Printf("Waiting %v before retry", sleep)
		time.Sleep(sleep)
	}
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&Receipt{}).Error; err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
