package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"golang.org/x/crypto/ed25519"
)

var cache [N]Message

// struct to store incoming messages
// type type.Message struct {
// 	Node      int    `json:"node"`
// 	Data      string `json:"data"`
// 	Sgnature []byte `json:"signature"`
// }

func sendMessage(conn *net.UDPConn, addr *net.UDPAddr, msg Message, publicKey ed25519.PublicKey) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
	}

	// Sign the message
	// signature := ed25519.Sign(privateKey, data)
	// msg.Signature = signature

	// Broadcast the message
	if _, err := conn.WriteToUDP(data, addr); err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
}

// func killPort(port int) {
// 	cmd := exec.Command("fuser", "-k", fmt.Sprintf("%d/udp", port))
// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Printf("Failed to kill port %d: %v", port, err)
// 	}
// }

func network2(node int, privateKey ed25519.PrivateKey, publicKeys [N]ed25519.PublicKey) {
	PORT := 10000 + node
	//killPort(PORT)
	conn, err2 := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: PORT})
	//_, err2 := reuseport.Listen("udp", ":"+strconv.Itoa(PORT))

	if err2 != nil {
		log.Fatalf("Failed to listen on UDP port: %v", err2)
	}

	defer conn.Close()

	for {
		b := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(b)
		if err != nil {
			log.Fatalf("Failed to read from UDP: %v", err)
		}

		var msg Message
		if err := json.Unmarshal(b[:n], &msg); err != nil {
			log.Fatalf("Failed to parse message: %v", err)
			m, _ := fmt.Printf("Node%d got a msg: %v", node, msg)
			fmt.Println(m)
		}

		// Verify the signature of the message
		if !ed25519.Verify(publicKeys[msg.Node], []byte(msg.Data), msg.Signature) {
			log.Fatalf("Invalid signature for message from node %d", msg.Node)
		}

		// Store the message in the cache
		cache[msg.Node] = msg

		time.Sleep(time.Second * 2)
		// Broadcast the message to other nodes
		for i, publicKey := range publicKeys {
			if i == node {
				continue
			}
			sendMessage(conn, addr, msg, publicKey)
		}
	}
}

func startNode2(node int, privateKey ed25519.PrivateKey, publicKeys [N]ed25519.PublicKey) {
	network2(node, privateKey, publicKeys)
}

func TestNodesSyncFromCache(t *testing.T) {
	privateKeys2 := [N]ed25519.PrivateKey{}
	publicKeys2 := [N]ed25519.PublicKey{}
	for i := 0; i < N; i++ {
		publicKey2, privateKey2, _ := ed25519.GenerateKey(rand.Reader)
		privateKeys2[i] = privateKey2
		publicKeys2[i] = publicKey2
	}

	// var wg sync.WaitGroup
	// wg.Add(N)

	for j := 0; j < N; j++ {
		go func(j int) {
			startNode2(j, privateKeys2[j], publicKeys2)
			//wg.Done()
		}(j)
	}

	// Wait for all nodes to start
	//wg.Wait()
	//time.Sleep(time.Second * 3)

	// Store data in cache for node 0
	data := "test data"
	signature := ed25519.Sign(privateKeys2[0], []byte(data))
	msg := Message{
		Node:      0,
		Data:      data,
		Signature: signature,
	}

	cache[0] = msg

	// Wait for data to be synced to other nodes
	time.Sleep(time.Second * 3)

	// Check if data is present in the cache for all nodes
	for i := 0; i < N; i++ {
		if cache[i].Data != data {
			t.Errorf("Data not synced to node %d", i)
		}
	}
}
