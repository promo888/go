package main

// import (
// 	"bytes"
// 	"crypto/sha256"
// 	"encoding/gob"
// 	"fmt"
// 	"log"

// 	"github.com/btcsuite/btcutil"
// )

// // MerkleTree represents a Merkle tree.
// type MerkleTree struct {
// 	RootNode *MerkleNode
// }

// // MerkleNode represents a node in a Merkle tree.
// type MerkleNode struct {
// 	Left  *MerkleNode
// 	Right *MerkleNode
// 	Hash  []byte
// }

// // NewMerkleTree creates a new Merkle tree from a slice of data.
// func NewMerkleTree(data [][]byte) *MerkleTree {
// 	var nodes []MerkleNode

// 	// create leaf nodes for the data
// 	for _, datum := range data {
// 		nodes = append(nodes, MerkleNode{nil, nil, sha256.Sum256(datum)})
// 	}

// 	// create parent nodes by combining adjacent nodes
// 	for len(nodes) > 1 {
// 		if len(nodes)%2 != 0 {
// 			nodes = append(nodes, nodes[len(nodes)-1])
// 		}
// 		var parentNodes []MerkleNode
// 		for i := 0; i < len(nodes); i += 2 {
// 			hash := sha256.Sum256(bytes.Join([][]byte{nodes[i].Hash, nodes[i+1].Hash}, []byte{}))
// 			parentNode := MerkleNode{&nodes[i], &nodes[i+1], hash[:]}
// 			parentNodes = append(parentNodes, parentNode)
// 		}
// 		nodes = parentNodes
// 	}

// 	return &MerkleTree{&nodes[0]}
// }


// // Block represents a block in the blockchain.
// type Block struct {
// 	Hash          []byte
// 	PrevHash      []byte
// 	Transactions  [][]byte
// 	MerkleTree    *MerkleTree
// 	TransactionsMerkleRoot []byte
// }

// // NewBlock creates a new block with the given transactions and previous block hash.
// func NewBlock(transactions [][]byte, prevHash []byte) *Block {
// 	// create a new Merkle tree for the transactions
// 	txTree := NewMerkleTree(transactions)

// 	// calculate the Merkle root hash for the transactions
// 	txRootHash := txTree.RootNode.Hash

// 	// create the block and calculate its hash
// 	block := &Block{
// 		PrevHash:      prevHash,
// 		Transactions:  transactions,
// 		MerkleTree:    txTree,
// 		TransactionsMerkleRoot: txRootHash,
// 	}
// 	block.Hash = calculateHash(block)

// 	return block
// }

// // Blockchain represents a blockchain.
// type Blockchain struct {
// 	blocks []*Block
// }

// // NewBlockchain creates a new blockchain with a genesis block.
// func NewBlockchain() *Blockchain {
// 	genesisBlock := NewBlock([][]byte{}, nil)
// 	return &Blockchain{
// 		blocks: []*Block{genesisBlock},
// 	}
// }

// // AddBlock adds a block to the blockchain.
// func (bc *Blockchain) AddBlock(transactions [][]byte) {
// 	prevBlock := bc.blocks[len(bc.blocks)-1]
// 	newBlock := NewBlock(transactions, prevBlock.Hash)

// 	bc.blocks = append(bc.blocks, newBlock)
// }

// // calculateHash calculates the hash for a block.
// func calculateHash(block *Block) []byte {
// 	var buf bytes.Buffer
// 	enc := gob.NewEncoder(&buf)
// 	err := enc.Encode(block)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	hash := sha256.Sum256(buf.Bytes())
// 	return hash[:]
// }

// func main() {
// 	// create a new blockchain with a genesis block
// 	bc := NewBlockchain()

// 	// add some blocks to the blockchain
// 	bc.AddBlock([][]byte{[]byte("transaction1"), []byte("transaction2")})
// 	bc.AddBlock([][]byte{[]byte("transaction3"), []byte("transaction4")})

// 	// print out the blockchain
// 	for _, block := range bc.blocks {
// 		fmt.Printf("Block %x\n", block.Hash)
// 		fmt.Printf("  PrevHash: %x\n", block.PrevHash)
// 		fmt.Printf("  Transactions: %v\n", block.Transactions)
// 		fmt.Printf("  TransactionsMerkleRoot: %x\n", block.TransactionsMerkleRoot)

// 		// verify the Merkle root hash for the block's transactions
// 		mt := btcutil.NewTxTree()
// 		for _, tx := range block.Transactions {
// 			mt.AddTx(btcutil.NewTx(tx))
// 		}
// 		if !bytes.Equal(mt.Hash(), block.TransactionsMerkleRoot) {
// 			fmt.Println("  TransactionsMerkleRoot verification failed")
// 		}

// 		// verify the Merkle root hash for the block itself
// 		blockData := bytes.Join([][]byte{block.PrevHash, block.TransactionsMerkleRoot}, []byte{})
// 		blockHash := sha256.Sum256(blockData)
// 		if !bytes.Equal(blockHash[:], block.Hash) {
// 			fmt.Println("  Block hash verification failed")
// 		}
// 	}
// }
