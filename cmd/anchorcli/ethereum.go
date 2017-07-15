package main

import (
	"context"
	"encoding/json"
	"fmt"
	blktk "github.com/Magicking/gethitihteg"
	ethtk "github.com/Magicking/gethitihteg/ethereum"
	"github.com/Magicking/rc-ge-ch-pdf/internal"
	"github.com/Magicking/rc-ge-ch-pdf/internal/merkle"
	"golang.org/x/crypto/sha3"
	"io"
	"log"
	"os"
)

const pkey = "18030537dbdd38d0764947d40bed98fc4d2a21af82765a7de7b13d2e4076773c"
const gethUrl = "http://1.1.3.7:8545/"

func sendData(data []byte) (string, error) {
	blkCtx, err := blktk.NewBlockchainContext(gethUrl, pkey)
	if err != nil {
		log.Fatalf("blktk.newblockchaincontext: %v", err)
	}
	anchor := ethtk.NewAnchor(&blkCtx.AO.Address, blkCtx.NC)
	tx, err := anchor.PrepareData(blkCtx.AO.Transactor, data)
	if err != nil {
		return "", fmt.Errorf("ethtk.NewAnchor: %v", err)
	}
	tx, err = blkCtx.AO.Sign(tx)
	if err != nil {
		return "", fmt.Errorf("blkCtx.AO.Sign: %v", err)
	}
	blkCtx.NC.SendTransaction(context.TODO(), tx)
	if err != nil {
		return "", fmt.Errorf("blkCtx.NC.SendTransaction: %v", err)
	}
	return tx.Hash().Hex(), nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s FILES...\n", os.Args[0])
	}
	hashs := make([]merkle.Hashable, len(os.Args)-1)
	//for each file in arguments
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		//open file
		f, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		h := sha3.New256()
		//sum256 the file
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		//associate filename with hash
		hash := h.Sum(nil)
		hashs[i-1] = internal.Hashitem(hash[:])
	}
	receipts, merkleRoot := internal.NewChainpoints(hashs)
	//send merkleroot
	txhash, err := sendData(merkleRoot)
	if err != nil {
		log.Fatalf("Could not send merkleRoot to Blockchain: %v", err)
	}
	for i, v := range receipts {
		//fill receipt
		v.Anchors = []internal.AnchorPoint{internal.AnchorPoint{SourceID: txhash, Type: "ETHData"}}
		receiptPath := fmt.Sprintf("%s.json", os.Args[i+1])
		//open file
		f, err := os.Create(receiptPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		enc := json.NewEncoder(f)
		err = enc.Encode(v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receipt %v written.\n", receiptPath)

	}
	//wait for transaction
	//write receipt
}
