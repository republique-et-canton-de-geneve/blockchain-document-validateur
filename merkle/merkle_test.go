package merkle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

func hop(left, right []byte) []byte {
	tmpcat := [][]byte{left, right}
	res := sha3.Sum256(bytes.Join(tmpcat, nil))
	return res[:]
}

func manual_verif() {
	H0 := common.Hex2Bytes("bdf8c9bdf076d6aff0292a1c9448691d2ae283f2ce41b045355e2c8cb8e85ef2")
	H1 := common.Hex2Bytes("2BAC4AD7D7C32B2C5D16A00ABBB1544E1D446541D24FA8B5925A26A21F73BEED")
	H2 := common.Hex2Bytes("C91DDF9D6598B2F3BD585ACEDF84333447251352A6E1EC446BF76427FE7CFF22")

	h0 := sha3.Sum256(H0)
	H0 = h0[:]
	//h1 := sha3.Sum256(H1)
	//H1 = h1[:]
	fmt.Println(common.Bytes2Hex(h0[:]))
	//fmt.Println(common.Bytes2Hex(h1[:]))
	H3 := hop(H0, H1)
	fmt.Println(common.Bytes2Hex(H3))
	H4 := hop(H3, H2)
	fmt.Println(common.Bytes2Hex(H4))
	//H3 := hop(H1, H2)
	//fmt.Println(common.Bytes2Hex(H3))
	//H4 := hop(H3, H1)
	//fmt.Println(common.Bytes2Hex(H4))
}

func testmerkle() {
	H0 := common.Hex2Bytes("bdf8c9bdf076d6aff0292a1c9448691d2ae283f2ce41b045355e2c8cb8e85ef3")
	H1 := common.Hex2Bytes("cb0dbbedb5ec5363e39be9fc43f56f321e1572cfcf304d26fc67cb6ea2e49faf")
	H2 := common.Hex2Bytes("da0ed1fecac504ea4f76d241a45032fa97b9eb692614419a04c9a9c32e39df2d")

	cpProofs := NewChainpoints([]Hashable{hashitem(H0), hashitem(H1), hashitem(H2)})
	rootH, proofs := ChainpointProofsFromHashables([]Hashable{hashitem(H0), hashitem(H1), hashitem(H2)})
	hash := SimpleHashFromTwoHashes(H0, nil)

	ok := Verify(hash, *proofs[0], rootH)
	if ok {
		fmt.Println("Okey !")
	} else {
		fmt.Println("Ko !")
	}
	for i := range cpProofs {
		ok = cpProofs[i].MerkleVerify()
		if ok {
			fmt.Println("Okey !")
		} else {
			fmt.Println("Ko !")
		}
	}
	fmt.Println(common.Bytes2Hex(rootH))
	proofjson, err := json.Marshal(proofs)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(string(proofjson))

	fmt.Println("----8<-------8<-------8<---")
	manual_verif()

	res := NewChainpoints([]Hashable{hashitem(H0), hashitem(H1), hashitem(H2)})
	proofjson, err = json.Marshal(res)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(string(proofjson))
}
