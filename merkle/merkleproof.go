package merkle

import (
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

type Hashitem []byte

func (hi Hashitem) Hash() []byte {
	ret := sha3.Sum256(hi)
	return ret[:]
}

func (hi Hashitem) Bytes() []byte {
	return []byte(hi)
}

type AnchorPoint struct {
	SourceID string `json:"sourceId"`
	Type     string `json:"type"`
}

type ChainpointLeafString struct {
	Left  string `json:"left,omitempty"`  // Hashes from leaf's sibling to a root's child.
	Right string `json:"right,omitempty"` // Hashes from leaf's sibling to a root's child.
}

type Chainpoint struct {
	Context    string                 `json:"@context"`
	Anchors    []AnchorPoint          `json:"anchors"`
	MerkleRoot string                 `json:"merkleRoot"`
	Proof      []ChainpointLeafString `json:"proof"`
	TargetHash string                 `json:"targetHash"`
	Type       string                 `json:"type"`
}

func NewChainpoints(items []Hashable) ([]Chainpoint, []byte) {
	rootH, proofs := ChainpointProofsFromHashables(items)
	if len(proofs) != len(items) {
		panic("Not all items were entered into merkle tree")
	}
	unanchoredReceipts := make([]Chainpoint, len(items))
	for i, v := range items {
		unanchoredReceipts[i].Type = "ChainpointSHA3-256v2"
		unanchoredReceipts[i].Context = "https://w3id.org/chainpoint/v2"
		unanchoredReceipts[i].Proof = proofs[i].Chainpoint()
		unanchoredReceipts[i].MerkleRoot = common.Bytes2Hex(rootH)
		unanchoredReceipts[i].TargetHash = common.Bytes2Hex(v.Bytes())
	}
	return unanchoredReceipts, rootH
}

func (cp *Chainpoint) MerkleVerify() bool {
	targetHash := common.Hex2Bytes(cp.TargetHash)
	aunts := ChainpointProofFromStringAunt(cp.Proof)
	root := common.Hex2Bytes(cp.MerkleRoot)
	return Verify(targetHash, aunts, root)
}
