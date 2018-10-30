/*
Computes a deterministic minimal height merkle tree hash.
If the number of items is not a power of two, some leaves
will be at different levels. Tries to keep both sides of
the tree the same size, but the left may be one greater.

Use this for short deterministic trees, such as the validator list.
For larger datasets, use IAVLTree.

                        *
                       / \
                     /     \
                   /         \
                 /             \
                *               *
               / \             / \
              /   \           /   \
             /     \         /     \
            *       *       *       h6
           / \     / \     / \
          h0  h1  h2  h3  h4  h5

*/

package merkle

import (
	"bytes"
	//"fmt"

	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/common"
)

type Hashable interface {
	Hash() []byte
	Bytes() []byte
}

// Chainpoint 2.1 Proof version
func SimpleHashFromTwoHashes(left []byte, right []byte) []byte {
	buffer := [][]byte{left, right}
	hashed := sha3.Sum256(bytes.Join(buffer, nil))

	return hashed[:]
}

func SimpleHashFromHashes(hashes [][]byte) []byte {
	// Recursive impl.
	switch len(hashes) {
	case 0:
		return nil
	case 1:
		return hashes[0]
	default:
		left := SimpleHashFromHashes(hashes[:(len(hashes)+1)/2])
		right := SimpleHashFromHashes(hashes[(len(hashes)+1)/2:])
		return SimpleHashFromTwoHashes(left, right)
	}
}

//--------------------------------------------------------------------------------

type ChainpointLeaf struct {
	Left  []byte // Hashes from leaf's sibling to a root's child.
	Right []byte // Hashes from leaf's sibling to a root's child.
}

type ChainpointProof struct {
	Aunts []ChainpointLeaf // Hashes from leaf's sibling to a root's child.
}

// proofs[0] is the proof for items[0].
func ChainpointProofsFromHashables(items []Hashable) (rootHash []byte, proofs []*ChainpointProof) {
	trails, rootSPN := trailsFromHashables(items)
	rootHash = rootSPN.Hash
	proofs = make([]*ChainpointProof, len(items))
	for i, trail := range trails {
		proofs[i] = &ChainpointProof{
			Aunts: trail.FlattenAunts(),
		}
	}
	return
}

func ChainpointProofFromStringAunt(aunts []ChainpointLeafString) ChainpointProof {
	_aunts := make([]ChainpointLeaf, len(aunts))
	for i, v := range aunts {
		var cpl ChainpointLeaf
		if v.Left != "" {
			cpl = ChainpointLeaf{Left: common.Hex2Bytes(v.Left)}
		} else if v.Right != "" {
			cpl = ChainpointLeaf{Right: common.Hex2Bytes(v.Right)}
		} else {
			panic("Left and right both unset !")
		}
		_aunts[i] = cpl
	}

	return ChainpointProof{Aunts: _aunts}
}

// Verify that leafHash is a leaf hash of the simple-merkle-tree
// which hashes to rootHash.
func Verify(targetHash []byte, proof ChainpointProof, rootHash []byte) bool {
	computedHash := computeHashFromAunts(targetHash, proof)
	if computedHash == nil {
		return false
	}
	if !bytes.Equal(computedHash, rootHash) {
		return false
	}
	return true
}

func (sp *ChainpointLeaf) Chainpoint() ChainpointLeafString {
	if sp.Left != nil {
		return ChainpointLeafString{Left: common.Bytes2Hex(sp.Left)}
	}
	return ChainpointLeafString{Right: common.Bytes2Hex(sp.Right)}
}

func (sp *ChainpointProof) Chainpoint() []ChainpointLeafString {
	var aunts []ChainpointLeafString

	for _, a := range sp.Aunts {
		aunts = append(aunts, a.Chainpoint())
	}

	return aunts
}

// Use the leafHash and innerHashes to get the root merkle hash.
// If the length of the innerHashes slice isn't exactly correct, the result is nil.
func computeHashFromAunts(targetHash []byte, proofs ChainpointProof) []byte {
	result := targetHash
	for _, proof := range proofs.Aunts {
		if proof.Left != nil {
			result = SimpleHashFromTwoHashes(proof.Left, result)
		} else if proof.Right != nil {
			result = SimpleHashFromTwoHashes(result, proof.Right)
		} else {
			panic("Left or right should be set for ChainpointProof")
		}
	}
	return result
}

// Helper structure to construct merkle proof.
// The node and the tree is thrown away afterwards.
// Exactly one of node.Left and node.Right is nil, unless node is the root, in which case both are nil.
// node.Parent.Hash = hash(node.Hash, node.Right.Hash) or
// 									  hash(node.Left.Hash, node.Hash), depending on whether node is a left/right child.
type SimpleProofNode struct {
	Hash   []byte
	Parent *SimpleProofNode
	Left   *SimpleProofNode // Left sibling  (only one of Left,Right is set)
	Right  *SimpleProofNode // Right sibling (only one of Left,Right is set)
}

// Starting from a leaf SimpleProofNode, FlattenAunts() will return
// the inner hashes for the item corresponding to the leaf.
func (spn *SimpleProofNode) FlattenAunts() (leaflist []ChainpointLeaf) {
	// Nonrecursive impl.
	for spn != nil {
		var lv ChainpointLeaf
		if spn.Left != nil {
			lv = ChainpointLeaf{Left: spn.Left.Hash}
		} else if spn.Right != nil {
			lv = ChainpointLeaf{Right: spn.Right.Hash}
		} else {
			break
		}
		spn = spn.Parent
		leaflist = append(leaflist, lv)
	}
	return leaflist
}

// trails[0].Hash is the leaf hash for items[0].
// trails[i].Parent.Parent....Parent == root for all i.
func trailsFromHashables(items []Hashable) (trails []*SimpleProofNode, root *SimpleProofNode) {
	// Recursive impl.
	switch len(items) {
	case 0:
		return nil, nil
	case 1:
		trail := &SimpleProofNode{items[0].Bytes(), nil, nil, nil}
		return []*SimpleProofNode{trail}, trail
	default:
		lefts, leftRoot := trailsFromHashables(items[:(len(items)+1)/2])
		rights, rightRoot := trailsFromHashables(items[(len(items)+1)/2:])
		rootHash := SimpleHashFromTwoHashes(leftRoot.Hash, rightRoot.Hash)
		root := &SimpleProofNode{rootHash, nil, nil, nil}
		leftRoot.Parent = root
		leftRoot.Right = rightRoot
		rightRoot.Parent = root
		rightRoot.Left = leftRoot
		return append(lefts, rights...), root
	}
}
