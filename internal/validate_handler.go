package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	blktk "github.com/Magicking/gethitihteg"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
	"io"
	"log"
	"net/http"
	"strings"
)

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

func validateReceipt(receipt *Chainpoint, hash common.Hash) error {
	if !receipt.MerkleVerify() {
		return fmt.Errorf("Invalid receipt")
	}
	targetHash := common.Hex2Bytes(receipt.TargetHash)

	if !bytes.Equal(targetHash, hash.Bytes()) {
		return fmt.Errorf("The hash in the receipt mismatch the file hash: %s != %s", common.Bytes2Hex(targetHash), common.Bytes2Hex(hash.Bytes()))
	}

	root := common.Hex2Bytes(receipt.MerkleRoot)

	var ok bool
	for _, v := range receipt.Anchors {
		if v.Type == "ETHData" {
			sourceId := common.HexToHash(v.SourceID)
			data, err := getData(sourceId)
			if err != nil {
				return fmt.Errorf("The transaction inexistent: %v\nTx hash: %s\n", err, v.SourceID)
			}
			if !bytes.Equal(data, root) {
				return fmt.Errorf("The receipt does not validate")
			}
			ok = true // IF ANY VALIDATE
		} else {
			log.Println("An unknown anchor type could be validated: ", v.Type)
		}
	}

	if !ok {
		log.Println("No anchor could be validated")
	}
	return nil
}

func ValidateHandler(ctx context.Context, prefix string, handler http.Handler) http.Handler {
	middle := func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, prefix) {
			handler.ServeHTTP(w, r)
			return
		}
		//parse the multipart form in the request
		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//get a ref to the parsed multipart form
		m := r.MultipartForm

		//get the *fileheaders
		files := m.File["myfiles"]
		if len(files) != 2 {
			http.Error(w, fmt.Sprintf("Invalid number of file, should be FILE + FILE RECEIPT"), http.StatusInternalServerError)
			return
		}
		receipt_found := false
		var hash []byte
		var receipt Chainpoint
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if !receipt_found {
				dec := json.NewDecoder(file)
				err = dec.Decode(&receipt)
			}
			if err == nil {
				receipt_found = true
				continue
			}
			h := sha3.New256()
			if _, err = io.Copy(h, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			hash = h.Sum(nil)
		}
		log.Println(receipt)
		if !receipt_found {
			http.Error(w, fmt.Sprintf("Invalid number of file, should be FILE + FILE RECEIPT"), http.StatusInternalServerError)
			return
		}
		if err = validateReceipt(&receipt, common.BytesToHash(hash)); err != nil {
			http.Error(w, fmt.Sprintf("Could not validate Receipt: %v", err), http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(middle)
}
