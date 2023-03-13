package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	//"runtime"
	//"sync"
	//"reuseport"

	n "xcoin.com/nodes"
	t "xcoin.com/v1/types"
)

func BenchmarkPost(b *testing.B) {
	//var mu sync.Mutex
	var cnt int
	port := "55555"
	go n.Start(port)

	//b.SetParallelism(10000)
	//runtime.GOMAXPROCS(1)
	b.ResetTimer()
	for i := 0; i < 1000; i++ {

		//new request
		//var msgBytes []byte
		newMsg, _ := t.NewTransaction("from", "to", "asset", 1.123, 0.001, "description")
		data, err := json.Marshal(newMsg)
		if err != nil {
			b.Fatalf("Failed to marshal data: %v", err) //todo ignore failure
		}
		resp, err := http.Post("http://localhost:"+port, "application/json", bytes.NewReader(data))
		if err != nil {
			b.Fatalf("Failed to send request: %v", err)
		} else {
			//mu.Lock()
			cnt++
			if cnt == 999 {
				fmt.Printf("%d - %d\n", cnt, resp.StatusCode) //, resp.Body
			}
			//defer mu.Unlock()
			//fmt.Println(resp.StatusCode) //, resp.Body)
		}

		if resp.StatusCode != http.StatusOK {
			b.Fatalf("Received non-200 response: %d", resp.StatusCode) ////todo ignore failure
		}
	}
}

func startSocketServer() {
	// Listen for incoming connections on port 11111
	listener, err := net.Listen("tcp", ":55555")
	//listener, err := reuseport.Listen("tcp", ":"+strconv.Itoa(tcpPort))
	if err != nil {
		log.Fatalf("\nFailed to start server: %v", err)
	}
	defer listener.Close()

	// Loop indefinitely, accepting incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("\nFailed to accept incoming connection: %v", err)
			continue
		}

		// set a deadline for the connection
		conn.SetDeadline(time.Now().Add(10 * time.Second))

		// Handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read data from the client
	buf := make([]byte, 1024) //1024 //320
	_, err := conn.Read(buf)  //n
	if err != nil {
		log.Printf("Failed to read data from client: %v", err)
		return
	}

	// Print the data received from the client
	//fmt.Printf("Received %d bytes from client: %s\n", n, string(buf[:n]))

	payload_size := 1000 //500000 //0
	// Send a response back to the client
	response := make([]byte, payload_size)          //"Hello, client!"
	if _, err := conn.Write(response); err != nil { //[]byte(response)
		log.Printf("Failed to send response to client: %v", err)
		return
	}
}

func TestSocketPayload(t *testing.T) {
	//start socket server
	go startSocketServer()
	// Create a connection to the server
	time.Sleep(100 * time.Millisecond)

	// conn, err := net.Dial("tcp", "localhost:55555")
	// if err != nil {
	// 	t.Fatalf("Failed to connect to server: %v", err)
	// }
	// defer conn.Close()

	// Create a 300-byte payload
	payload_size := 320 //300 //100000 //0 //000 //0
	payload := make([]byte, payload_size)

	// Measure the time it takes to send the payload to the server
	//b.ResetTimer()
	start := time.Now()
	iters := 7777
	for i := 0; i < iters; i++ {
		conn, err := net.Dial("tcp", "localhost:55555")
		if err != nil {
			t.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()
		//fmt.Println(i)
		if _, err := conn.Write(payload); err != nil {
			t.Fatalf("\nFailed to send data to server: %v", err)
		}
	}
	//b.StopTimer()
	elapsed := time.Since(start)
	fmt.Printf("%d requests with %d-byte payload took %s\n", iters, payload_size, elapsed)
}
