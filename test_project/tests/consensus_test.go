package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"strconv"
	"testing"
	"time"

	"golang.org/x/crypto/ed25519"
)

const (
	N         = 5 // number of nodes
	Buffer    = 1024
	TestData  = "test data"
	TestData2 = "test data 2"
)

// // struct to store incoming messages
type Message struct {
	Node      int    `json:"node"`
	Data      string `json:"data"`
	Signature []byte `json:"signature"`
}

func network(node int, privateKey ed25519.PrivateKey, publicKeys [N]ed25519.PublicKey) {
	port := strconv.Itoa(10000 + node)
	log.Println("Starting udp server on port:" + port)
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		log.Fatal("Failed to start udp server")
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	for {
		message := make([]byte, Buffer)
		_, from, err := conn.ReadFromUDP(message)
		if err != nil {
			log.Fatal(err)
			return
		}

		// Check for null bytes in the message
		// if bytes.Contains(message, []byte{0x00}) {
		// 	continue
		// }

		// Replace null bytes with empty string
		message = bytes.Replace(message, []byte{0x00}, []byte{}, -1)

		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Fatal(err)
			return
		}

		publicKey := publicKeys[msg.Node]
		if !ed25519.Verify(publicKey, []byte(msg.Data), msg.Signature) {
			continue
		}

		// respond to the message
		response, err := json.Marshal("OK")
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = conn.WriteToUDP(response, from)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func startNode(node int, privateKey ed25519.PrivateKey, publicKeys [N]ed25519.PublicKey) {
	go func() {
		network(node, privateKey, publicKeys)
	}()
}

// func startNode(node int, privateKey ed25519.PrivateKey, publicKeys [N]ed25519.PublicKey, t *testing.T) {
// 	go func() {
// 		network(node, privateKey, publicKeys, t)
// 	}()
// }

func TestConsensus(t *testing.T) {
	// generate private and public keys for each node
	var privateKeys [N]ed25519.PrivateKey
	var publicKeys [N]ed25519.PublicKey
	for i := 0; i < N; i++ {
		publicKey, privateKey, _ := ed25519.GenerateKey(nil)
		privateKeys[i] = privateKey
		publicKeys[i] = publicKey
		startNode(i, privateKey, publicKeys)
	}

	// send messages from each node
	for i := 0; i < N; i++ {
		data := []byte(TestData)
		signature := ed25519.Sign(privateKeys[i], data)

		message := Message{
			Node:      i,
			Data:      string(data),
			Signature: signature,
		}

		buffer, err := json.Marshal(message)
		if err != nil {
			t.Error(err)
			return
		}

		conn, err := net.Dial("udp", ":"+strconv.Itoa(10000+i))
		if err != nil {
			t.Error(err)
			return
		}

		_, err = conn.Write(buffer)
		if err != nil {
			t.Error(err)
			return
		}

		response := make([]byte, Buffer)
		_, err = conn.Read(response)
		if err != nil {
			t.Error(err)
			return
		}
	}

	// send an invalid message from a node //TestData2
	// data := []byte(TestData2)
	// signature := ed25519.Sign(privateKeys[0], data)

	// message := Message{
	// 	Node:      5,
	// 	Data:      string(data),
	// 	Signature: signature,
	// }

	// buffer, err := json.Marshal(message)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// conn, err := net.Dial("udp", ":10001")
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// _, err = conn.Write(buffer)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// check that the response is not received for the invalid message
	select {
	case <-time.After(time.Second):
		// expected behavior, do nothing
	default:
		t.Error("Response received for invalid message") //: " + message.Data)"
	}

}

// package main

// import (
// 	"encoding/json"
// 	"net"
// 	"testing"
// 	"time"

// 	"golang.org/x/crypto/ed25519"
// )

// const (
// 	N         = 5 // number of nodes
// 	Buffer    = 1024
// 	Port      = ":10001" // port for the UDP server in the test
// 	TestData  = "test data"
// 	TestData2 = "test data 2"
// )

// // struct to store incoming messages
// type Message struct {
// 	Node      int    `json:"node"`
// 	Data      string `json:"data"`
// 	Signature []byte `json:"signature"`
// }

// func TestConsensus(t *testing.T) {
// 	// generate private and public keys for each node
// 	var privateKeys [N]ed25519.PrivateKey
// 	var publicKeys [N]ed25519.PublicKey
// 	for i := 0; i < N; i++ {
// 		publicKey, privateKey, _ := ed25519.GenerateKey(nil)
// 		privateKeys[i] = privateKey
// 		publicKeys[i] = publicKey
// 	}

// 	// start the UDP server in a goroutine
// 	go main()

// 	// send messages from each node
// 	for i := 0; i < N; i++ {
// 		data := []byte(TestData)
// 		signature := ed25519.Sign(privateKeys[i], data)

// 		message := Message{
// 			Node:      i,
// 			Data:      string(data),
// 			Signature: signature,
// 		}

// 		buffer, err := json.Marshal(message)
// 		if err != nil {
// 			t.Error(err)
// 			return
// 		}

// 		conn, err := net.Dial("udp", Port)
// 		if err != nil {
// 			t.Error(err)
// 			return
// 		}

// 		_, err = conn.Write(buffer)
// 		if err != nil {
// 			t.Error(err)
// 			return
// 		}

// 		response := make([]byte, Buffer)
// 		_, err = conn.Read(response)
// 		if err != nil {
// 			t.Error(err)
// 			return
// 		}
// 	}

// 	// send an invalid message from a node
// 	data := []byte(TestData2)
// 	signature := ed25519.Sign(privateKeys[0], data)

// 	message := Message{
// 		Node:      5,
// 		Data:      string(data),
// 		Signature: signature,
// 	}

// 	buffer, err := json.Marshal(message)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	conn, err := net.Dial("udp", Port)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	_, err = conn.Write(buffer)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	// check that the response is not received for the invalid message
// 	select {
// 	case <-time.After(time.Second):
// 		// expected behavior, do nothing
// 	default:
// 		t.Error("Response received for invalid message")
// 	}
// }
