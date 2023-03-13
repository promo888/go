package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"reflect"
	"sync"
	"testing"
	//"time"
)



// Node struct
type Node struct {
	ID          int
	blockchain  []Block
	transaction chan Transaction
	mu          sync.Mutex
}

// Transaction struct
type Transaction struct {
	Sender   string
	Receiver string
	Amount   int
	BlockNum int //todo big int
}

// Block struct
type Block struct {
	ID           int
	Transactions []Transaction
	PrevHash     string
	Hash         string
}

// RPC request struct
type RPCRequest struct {
	NodeID      int
	Blockchain  []Block
	Transaction Transaction
}

// RPC response struct
type RPCResponse struct {
	OK      bool
	Message string
}

// VerifyConsensus RPC function
func (n *Node) VerifyConsensus(request RPCRequest, response *RPCResponse) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	// Check that the requesting node's blockchain matches the other nodes' blockchains
	for i:= range request.Blockchain {
		if len(nodes[i].blockchain) < len(n.blockchain) || !blockchainsMatch(n.blockchain, nodes[i].blockchain) {
			response.OK = false
			response.Message = "Blockchains do not match"
			return nil
		}
	}

	// Add the transaction to the requesting node's transaction channel
	n.transaction <- request.Transaction

	response.OK = true
	response.Message = fmt.Sprintf("Transaction accepted: %v", request.Transaction)

	return nil
}

func blockchainsMatch(a, b []Block) bool {
	if len(a) != len(b) {
		return false
	}
	for i, block := range a {
		if !reflect.DeepEqual(block, b[i]) {
			return false
		}
	}
	return true
}

// VerifyConsensus RPC function
// func VerifyConsensus(request RPCRequest, response *RPCResponse) error {
// 	// Check that the requesting node's blockchain matches the other nodes' blockchains
// 	for i, b := range request.Blockchain {
// 		if !reflect.DeepEqual(b, nodes[i].blockchain[len(nodes[i].blockchain)-1]) {
// 			response.OK = false
// 			response.Message = "Blockchains do not match"
// 			return nil
// 		}
// 	}

// 	// Add the transaction to the requesting node's transaction channel
// 	nodes[request.NodeID].transaction <- request.Transaction

// 	response.OK = true
// 	response.Message = "Transaction accepted"

// 	return nil
// }

// VerifyConsensus RPC function
func VerifyConsensus(request RPCRequest, response *RPCResponse) error {
    // Check that the requesting node's blockchain matches the other nodes' blockchains
    for i, b := range request.Blockchain {
        if !reflect.DeepEqual(b, nodes[i].blockchain[len(nodes[i].blockchain)-1]) {
            response.OK = false
            response.Message = "Blockchains do not match"
            return nil
        }
    }

    // Add the transaction to the requesting node's transaction channel
    for i := range nodes {
        if i == request.NodeID {
            continue
        }
        nodes[i].transaction <- request.Transaction
    }

    // Update block number on transaction
    tx := request.Transaction
    tx.BlockNum = len(nodes[request.NodeID].blockchain) - 1

    // Append transaction to the requesting node's blockchain
    nodes[request.NodeID].blockchain[len(nodes[request.NodeID].blockchain)-1].Transactions = append(
        nodes[request.NodeID].blockchain[len(nodes[request.NodeID].blockchain)-1].Transactions,
        tx,
    )

    response.OK = true
    response.Message = "Transaction accepted"

    return nil
}



func fanoutTransaction(tx Transaction) {
	var wg sync.WaitGroup
	wg.Add(len(nodes))
	for i, node := range nodes {
		go func(node *Node, i int) {
			req := RPCRequest{
				NodeID:      i,
				Blockchain:  node.blockchain,
				Transaction: tx,
			}
			var res RPCResponse
			client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%d", portnum+i))
			if err != nil {
				log.Fatal("Error dialing:", err)
			}
			err = client.Call("Node.VerifyConsensus", req, &res)
			if err != nil {
				log.Fatal("Error calling VerifyConsensus:", err)
			}
			if !res.OK {
				fmt.Printf("Node %d: %s", node.ID, res.Message)
			}
			wg.Done()
		}(node, i)

    }
}

// Nodes
var nodes []*Node
var portnum int = 10001
func TestSyncNodes(t *testing.T) {
	// Initialize nodes
	nodes = []*Node{}
	for i := 0; i < 10; i++ {
		nodes = append(nodes, &Node{ID: i, blockchain: []Block{{ID: 0}}, transaction: make(chan Transaction)})
	}


	// Start RPC server on each node
	for i := range nodes {
		rpc.Register(nodes[i])
	}

	// Start HTTP server on each node
	for i, n := range nodes {
        rpcServer := rpc.NewServer()
        rpcServer.Register(n)
        rpcServer.HandleHTTP(rpc.DefaultRPCPath+fmt.Sprintf("%d", portnum+i), rpc.DefaultDebugPath+fmt.Sprintf("%d", portnum+i))
        l, err := net.Listen("tcp", ":"+  fmt.Sprintf("%d", portnum+i))
        if err != nil {
            log.Fatal("Error listening:", err)
        }
        go func() {
            if err := http.Serve(l, nil); err != nil {
                log.Fatal("Error serving:", err)
            }
        }()
    }


	// //Create a new transaction and broadcast it to all nodes
	tx := Transaction{Sender:  "Alice", Receiver: "Bob", Amount: 10}
	tx1 := Transaction{Sender: "Charlie", Receiver: "Bob", Amount: 10}
	tx2 := Transaction{Sender: "Bob", Receiver: "Eva", Amount: 10}
	// var wg sync.WaitGroup
	// wg.Add(len(nodes))
	// for i, node := range nodes {
	// 	go func(node *Node, i int) {
	// 		req := RPCRequest{
	// 			NodeID:      i,
	// 			Blockchain:  node.blockchain,
	// 			Transaction: tx,
	// 		}
	// 		var res RPCResponse
	// 		client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%d", portnum+i))
	// 		if err != nil {
	// 			log.Fatal("Error dialing:", err)
	// 		}
	// 		err = client.Call("Node.VerifyConsensus", req, &res)
	// 		if err != nil {
	// 			log.Fatal("Error calling VerifyConsensus:", err)
	// 		}
	// 		if !res.OK {
	// 			t.Errorf("Node %d: %s", node.ID, res.Message)
	// 		}
	// 		wg.Done()
	// 	}(node, i)


    // }
	fanoutTransaction(tx)
	fanoutTransaction(tx1)
	fanoutTransaction(tx2)
    



	// //time.Sleep(1 * time.Second)
	// for _, n := range nodes {
    //     fmt.Printf("Node %d: blockchain: %v, transaction: %v\n", n.ID, n.blockchain, n.transaction)
    // }  

	// Create a new transaction and broadcast it to all nodes
	// tx := Transaction{Sender: "Alice", Receiver: "Bob", Amount: 10}
	// for i := range nodes {
	// 	req := RPCRequest{
	// 		NodeID:      i,
	// 		Blockchain:  nodes[i].blockchain,
	// 		Transaction: tx,
	// 	}
	// 	var res RPCResponse
	// 	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%d", portnum+i))
	// 	if err != nil {
	// 		log.Fatal("Error dialing:", err)
	// 	}
	// 	err = client.Call("Node.VerifyConsensus", req, &res)
	// 	if err != nil {
	// 		log.Fatal("Error calling VerifyConsensus:", err)
	// 	}
	// 	fmt.Println("Response: ",res.Message)
	// }

	// Wait for transactions to be processed
	for i := range nodes {
		for {
			select {
			case tx := <-nodes[i].transaction:
				nodes[i].mu.Lock()
				nodes[i].blockchain[len(nodes[i].blockchain)-1].Transactions = append(
					nodes[i].blockchain[len(nodes[i].blockchain)-1].Transactions,
					tx,
				)
				nodes[i].mu.Unlock()
			default:
				goto done
			}
		}
	}
done:

	// Print the final state of each node's blockchain with consensus state and the hash
	for i, n := range nodes {
		fmt.Printf("Node %d\n", i)
		for _, b := range n.blockchain {
			fmt.Printf("\tBlock %d:\n", b.ID)
		for _, t := range b.Transactions {
			fmt.Printf("\t\tTransaction from %s to %s for %d\n", t.Sender, t.Receiver, t.Amount)
		}
		fmt.Printf("\t\tPrevHash: %s\n", b.PrevHash)
		fmt.Printf("\t\tHash: %s\n", b.Hash)
		}
		//fmt.Printf("\tConsensus: %t\n", VerifyConsensus(n.blockchain))
		fmt.Println()
		}
		
		
	
}