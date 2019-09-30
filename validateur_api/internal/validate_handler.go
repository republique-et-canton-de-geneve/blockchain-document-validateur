package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"

	//	"os"
	"strings"

	rcge "github.com/geneva_validateur/merkle"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/crypto/sha3"
)

var BEG_DELIM = []byte(":JsOnbegin:")
var END_DELIM = []byte(":JsOnend:")

func isProtectedV(V *big.Int) bool {
	if V.BitLen() <= 8 {
		v := V.Uint64()
		return v != 27 && v != 28
	}
	// anything not 27 or 28 are considered unprotected
	return true
}

// deriveChainId derives the chain id from the given v parameter
func deriveChainId(v *big.Int) *big.Int {
	if v.BitLen() <= 64 {
		v := v.Uint64()
		if v == 27 || v == 28 {
			return new(big.Int)
		}
		return new(big.Int).SetUint64((v - 35) / 2)
	}
	v = new(big.Int).Sub(v, big.NewInt(35))
	return v.Div(v, big.NewInt(2))
}

func deriveSigner(V *big.Int) types.Signer {
	if V.Sign() != 0 && isProtectedV(V) {
		return types.NewEIP155Signer(deriveChainId(V))
	} else {
		return types.HomesteadSigner{}
	}
}

func getData(ctx context.Context, hash common.Hash) ([]byte, *big.Int, string, error) {
	ccCtx, ok := CCFromContext(ctx)
	if !ok {
		log.Fatalf("Could not obtain ClientConnector from context\n")
	}

	tx, hdr_hash, err := ccCtx.TransactionByHashFull(context.TODO(), hash)
	log.Println(tx, hdr_hash, err)
	if err != nil {
		return nil, nil, "", fmt.Errorf("ccCtx.TransactionByHashFull: %v\n", err)
	}
	if hdr_hash == nil {
		return nil, nil, "", fmt.Errorf("Transaction pending")
	}
	hdr, err := ccCtx.HeaderByHash(context.TODO(), *hdr_hash)
	if err != nil {
		return nil, nil, "", fmt.Errorf("ccCtx.TransactionByHashFull: %v\n", err)
	}
	v, _, _ := tx.RawSignatureValues()
	var from string
	if v != nil {
		// make a best guess about the signer and use that to derive
		// the sender.
		signer := deriveSigner(v)
		if f, err := types.Sender(signer, tx); err != nil { // derive but don't cache
			from = "[invalid sender: invalid sig]"
		} else {
			from = fmt.Sprintf("%x", f[:])
		}
	} else {
		from = "[invalid sender: nil V field]"
	}
	return tx.Data(), hdr.Time, from, nil
}

func validateReceipt(ctx context.Context, receipt *rcge.Chainpoint, hash common.Hash) (*big.Int, string, error) {
	if !receipt.MerkleVerify() {
		return nil, "", fmt.Errorf("Invalid receipt")
	}
	targetHash := common.Hex2Bytes(receipt.TargetHash)

	if !bytes.Equal(targetHash, hash.Bytes()) {
		return nil, "", fmt.Errorf("The hash in the receipt mismatch the file hash: %s != %s", common.Bytes2Hex(targetHash), common.Bytes2Hex(hash.Bytes()))
	}

	root := common.Hex2Bytes(receipt.MerkleRoot)

	for _, v := range receipt.Anchors {
		if v.Type == "ETHData" {
			sourceId := common.HexToHash(v.SourceID)
			log.Println(ctx, sourceId)
			data, anchor_date, from, err := getData(ctx, sourceId)
			log.Println(data, anchor_date, from, err)
			log.Printf("anchor_date: %v, from: %v", anchor_date, from)
			if err != nil {
				return nil, "", fmt.Errorf("The transaction inexistent: %v\nTx hash: %s\n", err, v.SourceID)
			}
			if !bytes.Equal(data, root) {
				return nil, "", fmt.Errorf("The receipt does not validate")
			}
			// IF ANY VALIDATE
			return anchor_date, from, nil
		} else {
			log.Println("An unknown anchor type could be validated: ", v.Type)
		}
	}
	return nil, "", fmt.Errorf("No anchor could be validated")
}

type ValidateResponse struct {
	TargetHash string   `json:"target_hash"`
	From       string   `json:"from"`
	Time       *big.Int `json:"time"`
}

func ExtractJson(r io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	begin := bytes.Index(data, BEG_DELIM)
	end := bytes.Index(data, END_DELIM)
	if begin < 0 || end < 0 {
		return nil, fmt.Errorf("File is not a receipt")
	}
	data = data[begin+len(BEG_DELIM) : end]
	data = bytes.Replace(data, []byte("\\173"), []byte("{"), -1)
	data = bytes.Replace(data, []byte("\\175"), []byte("}"), -1)
	return data, nil
}

func ValidateHandler(ctx context.Context, prefix, lockedAddress string, handler http.Handler) http.Handler {
	middle := func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, prefix) {
			handler.ServeHTTP(w, r)
			return
		}
		//parse the multipart form in the request
		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
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
		var receipt rcge.Chainpoint
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			if !receipt_found {
				data, err := ExtractJson(file)
				if err == nil {
					err = json.Unmarshal(data, &receipt)
					if err == nil {
						receipt_found = true
						continue
					}
				}
			}
			h := sha3.New256()
			if _, err = file.Seek(0, 0); err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			if _, err = io.Copy(h, file); err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			hash = h.Sum(nil)
		}
		if !receipt_found {
			http.Error(w, fmt.Sprintf("Invalid receipt file, should be FILE + FILE RECEIPT"), http.StatusInternalServerError)
			return
		}
		anchor_date, from, err := validateReceipt(ctx, &receipt, common.BytesToHash(hash))
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not validate Receipt: %v", err), http.StatusInternalServerError)
			return
		}

		var fromIsValid bool
		fromIsValid = false

		pubKeys := os.Getenv("LOCKED_ADDR")
		pubKeysArr := strings.Split(pubKeys, ",")

		timestampPubKeys := os.Getenv("TIMEOUT")
		timestampArr := strings.Split(timestampPubKeys, ",")

		for i, address := range pubKeysArr {
			if strings.TrimPrefix(address, "0x") == from {
				number, _ := strconv.Atoi(timestampArr[i])
				timeout := big.NewInt(int64(number))

				// if anchor_date > expiration time then not valid
				if timeout.Cmp(anchor_date) == -1 {
					fromIsValid = false
					break
				}
				fromIsValid = true
				break
			}
		}

		//if from != lockedAddress {
		if !fromIsValid {
			http.Error(w, fmt.Sprintf("Invalid receipt, could not valid submitter"), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		enc := json.NewEncoder(w)
		err = enc.Encode(&ValidateResponse{
			TargetHash: common.Bytes2Hex(hash),
			From:       from,
			Time:       anchor_date,
		})
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(middle)
}
