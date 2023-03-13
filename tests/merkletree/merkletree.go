package main

// import (
// 	"crypto/sha256"
// 	"fmt"
// )

// type Transaction struct {
// 	Sender   string
// 	Receiver string
// 	Amount   float64
// }

// type Block struct {
// 	PrevHash     []byte
// 	Transactions []*Transaction
// }

// func hashTransactions(transactions []*Transaction) []byte {
// 	var hashes [][]byte
// 	for _, t := range transactions {
// 		hashes = append(hashes, hashTransaction(t))
// 	}
// 	return hashHashes(hashes)
// }

// func hashTransaction(transaction *Transaction) []byte {
// 	data := []byte(fmt.Sprintf("%s%s%f", transaction.Sender, transaction.Receiver, transaction.Amount))
// 	hash := sha256.Sum256(data)
// 	hashSlice := make([]byte, 32)
// 	copy(hashSlice, hash[:])
// 	return hashSlice
// }

// func hashHashes(hashes [][]byte) []byte {
// 	if len(hashes) == 0 {
// 		return nil
// 	} else if len(hashes) == 1 {
// 		return hashes[0]
// 	} else if len(hashes)%2 == 1 {
// 		hashes = append(hashes, hashes[len(hashes)-1])
// 	}
// 	var parentHashes [][]byte
// 	for i := 0; i < len(hashes); i += 2 {
// 		hash := sha256.Sum256(append(hashes[i], hashes[i+1]...))
// 		parentHashes = append(parentHashes, make([]byte, 32))
// 		copy(parentHashes[len(parentHashes)-1], hash[:])
// 	}
// 	return hashHashes(parentHashes)
// }

// func main() {
// 	// Example input data.
// 	t1 := &Transaction{Sender: "Alice", Receiver: "Bob", Amount: 1.23}
// 	t2 := &Transaction{Sender: "Bob", Receiver: "Charlie", Amount: 4.56}
// 	t3 := &Transaction{Sender: "Charlie", Receiver: "Dave", Amount: 7.89}
// 	block1 := &Block{PrevHash: nil, Transactions: []*Transaction{t1, t2}}
// 	block2 := &Block{PrevHash: hashTransactions(block1.Transactions), Transactions: []*Transaction{t3}}
// 	//block3 := &Block{PrevHash: hashTransactions(block2.Transactions), Transactions: []*Transaction{}}
// 	transactions := []*Transaction{t1, t2, t3}

// 	// Compute the Merkle root of the transactions in the blocks.
// 	block1Hash := hashTransactions(block1.Transactions)
// 	block2Hash := hashTransactions(block2.Transactions)
// 	//block3Hash := hashTransactions(block3.Transactions)
// 	rootHash := hashHashes([][]byte{block1Hash, block2Hash}) //, block3Hash})

// 	// Verify that the computed Merkle root matches the expected value.
// 	expectedRootHash := hashHashes([][]byte{hashTransactions(transactions)})
// 	if string(rootHash) == string(expectedRootHash) {
// 		fmt.Println("Merkle root is valid!")
// 	} else {
// 		fmt.Println("Merkle root is not valid!")
// 	}
// }



// import (
// 	"crypto/sha256"
// 	"fmt"
// )

// type MerkleTree struct {
// 	leaves [][]byte
// }

// func NewMerkleTree(leaves [][]byte) *MerkleTree {
// 	return &MerkleTree{leaves}
// }

// func (t *MerkleTree) Root() []byte {
// 	hashes := t.leaves
// 	for len(hashes) > 1 {
// 		if len(hashes)%2 == 1 {
// 			hashes = append(hashes, hashes[len(hashes)-1])
// 		}
// 		var parentHashes [][]byte
// 		for i := 0; i < len(hashes); i += 2 {
// 			parentHashes = append(parentHashes, hashHashes(hashes[i], hashes[i+1]))
// 		}
// 		hashes = parentHashes
// 	}
// 	return hashes[0]
// }

// func (t *MerkleTree) VerifyRoot(leaf []byte, root []byte) bool {
// 	hashes := t.leaves
// 	for len(hashes) > 1 {
// 		if len(hashes)%2 == 1 {
// 			hashes = append(hashes, hashes[len(hashes)-1])
// 		}
// 		var parentHashes [][]byte
// 		for i := 0; i < len(hashes); i += 2 {
// 			parentHashes = append(parentHashes, hashHashes(hashes[i], hashes[i+1]))
// 		}
// 		hashes = parentHashes
// 	}
// 	return (len(hashes) == 1) && (hashHashes(hashes[0], leaf) == root)
// }

// func hashHashes(a []byte, b []byte) []byte {
// 	return sha256.Sum256(append(a, b...))[:]
// }

// func main() {
// 	// Create some example data.
// 	data := [][]byte{
// 		[]byte("foo"),
// 		[]byte("bar"),
// 		[]byte("baz"),
// 	}

// 	// Create a Merkle tree and compute its root.
// 	tree := NewMerkleTree(data)
// 	root := tree.Root()

// 	// Verify that the root is valid using a leaf node.
// 	leaf := data[1]
// 	if tree.VerifyRoot(leaf, root) {
// 		fmt.Println("Merkle root is valid!")
// 	} else {
// 		fmt.Println("Merkle root is not valid!")
// 	}
// }



// import (
// 	"crypto/sha256"
// 	"fmt"
// )

// // Node represents a node in the Merkle tree.
// type Node struct {
// 	Left  *Node
// 	Right *Node
// 	Hash  []byte
// }

// // BuildTree builds a Merkle tree from the provided data.
// func BuildTree(data [][]byte) *Node {
// 	// If the data has an odd number of elements, duplicate the last element.
// 	if len(data)%2 != 0 {
// 		data = append(data, data[len(data)-1])
// 	}

// 	// Build the tree bottom-up, starting with the leaf nodes.
// 	var nodes []*Node
// 	for _, d := range data {
// 		hash := sha256.Sum256(d)
// 		nodes = append(nodes, &Node{Hash: hash[:]})
// 	}

// 	// Merge pairs of nodes into parent nodes until there is only one node left.
// 	for len(nodes) > 1 {
// 		var parents []*Node
// 		for i := 0; i < len(nodes); i += 2 {
// 			left := nodes[i]
// 			right := nodes[i+1]
// 			hash := sha256.Sum256(append(left.Hash, right.Hash...))
// 			parents = append(parents, &Node{Left: left, Right: right, Hash: hash[:]})
// 		}
// 		nodes = parents
// 	}

// 	return nodes[0]
// }

// // VerifyProof verifies that the hash at the given index in the data array
// // is included in the Merkle tree with the given root and proof.
// func VerifyProof(root, hash []byte, index int, proof [][]byte) bool {
// 	// Start with the hash at the given index.
// 	h := hash
// 	//var data [32]byte

// 	// Verify each proof element.
// 	for _, p := range proof {
// 		// If the index is even, the current hash is on the left.
		
// 		if index%2 == 0 {		

// 			h = sha256.Sum256(append(h, p...))[:]
			
// 		} else {
// 			h = sha256.Sum256(append(p, h...))[:]
// 		}
// 		//copy(data[:], h[:])

// 		// Move up one level in the tree.
// 		index /= 2
// 	}

// 	// The resulting hash should match the root hash.
// 	return fmt.Sprintf("%x", h) == fmt.Sprintf("%x", root)
// }

// func main() {
// 	// Example usage.
// 	data := [][]byte{
// 		[]byte("hello"),
// 		[]byte("world"),
// 		[]byte("how"),
// 		[]byte("are"),
// 		[]byte("you"),
// 	}
// 	tree := BuildTree(data)
// 	fmt.Printf("Root hash: %x\n", tree.Hash)

// 	// Verify a proof for the hash at index 1.
// 	proof := [][]byte{tree.Right.Hash}
// 	if VerifyProof(tree.Hash, data[1], 1, proof) {
// 		fmt.Println("Proof is valid!")
// 	} else {
// 		fmt.Println("Proof is not valid!")
// 	}
// }
