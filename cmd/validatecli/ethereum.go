package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	blktk "github.com/Magicking/gethitihteg"
	"github.com/Magicking/rc-ge-ch-pdf/internal"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
	"io"
	"log"
	"os"
)

const gethUrl = "http://localhost:8545/"

func getData(hash common.Hash) ([]byte, error) {
	blkNC, err := blktk.NewNodeConnector(gethUrl, 3)
	if err != nil {
		log.Fatalf("blktk.newblockchaincontext: %v\n", err)
	}
	tx, isPending, err := blkNC.GetTransaction(context.TODO(), hash)
	if err != nil {
		return nil, fmt.Errorf("blkNC.GetTransaction: %v\n", err)
	}
	if isPending {
		log.Println("Pending transaction")
	}
	return tx.Data(), nil
}

func main() {
	if len(os.Args) != 3 {
		panic(fmt.Sprintf("Usage: %s FILE RECEIPT_FILE", os.Args[0]))
	}
	//open file
	f, err := os.Open(os.Args[1])
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
	f, err = os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	var receipt internal.Chainpoint
	err = dec.Decode(&receipt)
	if err != nil {
		log.Fatal(err)
	}

	if !receipt.MerkleVerify() {
		log.Fatalf("Invalid receipt")
	}
	targetHash := common.Hex2Bytes(receipt.TargetHash)

	if !bytes.Equal(targetHash, hash) {
		log.Fatalf("The file mismatch the file.\nFile: %s\nReceipt: %s\n", os.Args[1], os.Args[2])
	}

	root := common.Hex2Bytes(receipt.MerkleRoot)

	var ok bool
	for _, v := range receipt.Anchors {
		if v.Type == "ETHData" {
			sourceId := common.HexToHash(v.SourceID)
			data, err := getData(sourceId)
			if err != nil {
				log.Fatalf("The transaction inexistent: %v\nTx hash: %s\n", err, v.SourceID)
			}
			if !bytes.Equal(data, root) {
				log.Fatalf("The receipt does not validate .\nFile: %s\nReceipt: %s\n", os.Args[1], os.Args[2])
			}
			ok = true // IF ANY VALIDATE
		} else {
			log.Println("An unknown anchor type could be validated: ", v.Type)
		}
	}

	if !ok {
		log.Println("No anchor could be validated")
	}
	// Hash it (FileHash)
	// Open Json receipt
	// Verify targetHash == FileHash or return false
	// Get Transaction Data
	// Verify merkleRoot == Transaction Data
	// Return verify Merkle Tree
	log.Printf("File %s exists at least since TODO", os.Args[1])
}
