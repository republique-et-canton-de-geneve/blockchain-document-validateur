package internal

import (
	"context"
	"fmt"
	ethtk "github.com/Magicking/gethitihteg/ethereum"
	"github.com/Magicking/rc-ge-ch-pdf/internal/merkle"
	"golang.org/x/crypto/sha3"
	"io"
	"log"
	"net/http"
	"strings"
)

func sendData(ctx context.Context, data []byte) (string, error) {
	blkCtx, ok := BLKFromContext(ctx)
	if !ok {
		log.Fatalf("blktk.newblockchaincontext")
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

func UploadHandler(ctx context.Context, prefix string, handler http.Handler) http.Handler {
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
		hashs := make([]merkle.Hashable, len(files))
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			h := sha3.New256()
			//sum256 the file
			if _, err := io.Copy(h, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//associate filename with hash
			hash := h.Sum(nil)
			hashs[i] = Hashitem(hash[:])
			/* //copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
			}*/
		}
		if len(files) == 1 {
			hashs = append(hashs, hashs[0])
		}
		receipts, merkleRoot := NewChainpoints(hashs)
		//send merkleroot
		txhash, err := sendData(ctx, merkleRoot)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for i, v := range receipts {
			if len(files) == 1 && i > 0 {
				break
			}
			//fill receipt
			v.Anchors = []AnchorPoint{AnchorPoint{SourceID: txhash, Type: "ETHData"}}
			InsertReceipt(ctx, files[i].Filename, &v)
		}
		fmt.Println(merkleRoot)
		fmt.Println(receipts)
	}

	return http.HandlerFunc(middle)
}
