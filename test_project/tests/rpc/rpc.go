package main

// import (
// 	"fmt"
// 	"net"
// 	"net/rpc"
// 	"time"

// 	"github.com/syndtr/goleveldb/leveldb"
// 	//"github.com/syndtr/goleveldb/leveldb/util"
// 	//"github.com/vmihailenco/msgpack/v4"
// 	"encoding/json"
// )

// type Payload struct {
// 	Type string          `json:"type"`
// 	Data json.RawMessage  `json:"data"` //msgpack.RawMessage
// }

// type Person struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// type Animal struct {
// 	Name  string `json:"name"`
// 	Sound string `json:"sound"`
// }

// type PayloadResponse struct {
// 	Data []byte
// }

// type PayloadService struct{}

// func (s *PayloadService) Process(payload Payload, response *PayloadResponse) error {
// 	// store payload in LevelDB
// 	db, err := leveldb.OpenFile("payloads.db", nil)
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	payloadData, err := json.Marshal(payload)
// 	if err != nil {
// 		return err
// 	}
// 	err = db.Put([]byte(time.Now().String()), payloadData, nil)
// 	if err != nil {
// 		return err
// 	}

// 	// parse payload into specific objects
// 	var data interface{}
// 	switch payload.Type {
// 	case "person":
// 		var person Person
// 		err := json.Unmarshal(payload.Data, &person)
// 		if err != nil {
// 			return err
// 		}
// 		data = person
// 	case "animal":
// 		var animal Animal
// 		err := json.Unmarshal(payload.Data, &animal)
// 		if err != nil {
// 			return err
// 		}
// 		data = animal
// 	default:
// 		return fmt.Errorf("unknown payload type: %s", payload.Type)
// 	}

// 	// set response data
// 	responseData, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}
// 	response.Data = responseData
// 	return nil
// }

// func main() {
// 	// create RPC server
// 	payloadService := new(PayloadService)
// 	rpc.Register(payloadService)
// 	listener, err := net.Listen("tcp", ":11111")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// start server in a goroutine
// 	go rpc.Accept(listener)

// 	// create RPC client
// 	client, err := rpc.Dial("tcp", "localhost:11111")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// benchmark with person payload
// 	personPayload := Payload{
// 		Type: "person",
// 		Data:[]byte(`{"name":"John","age":30}`), //msgpack.RawMessage([]byte(`{"name":"John","age":30}`) ),
// 	}
// 	start := time.Now()
// 	for i := 0; i < 10000; i++ {
// 		var response PayloadResponse
// 		err = client.Call("PayloadService.Process", personPayload, &response)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	elapsed := time.Since(start)
// 	fmt.Printf("Person payload took %s\n", elapsed)

// 	// benchmark with animal payload
// 	animalPayload := Payload{
// 		Type: "animal",
// 		Data:[]byte(`{"name":"Dog","sound":"Bark"}`), //msgpack.RawMessage([]byte(`{"name":"Dog","sound":"Bark"}`)),
// 	}
// 	start = time.Now()
// 	for i := 0; i < 1; i++ {
// 		var response PayloadResponse
// 		err = client.Call("PayloadService.Process", animalPayload, &response)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// 	elapsed = time.Since(start)
// 	fmt.Printf("Animal payload took %s\n, ", elapsed)

// }

import (
	"fmt"
	"net"
	"net/rpc"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	//"github.com/syndtr/goleveldb/leveldb/util"
	"testing"
)

type Payload struct {
	Data []byte
}

type PayloadResponse struct {
	Data []byte
}

type PayloadService struct{}

func (s *PayloadService) Process(payload Payload, response *PayloadResponse) error {
	// store payload in LevelDB
	db, err := leveldb.OpenFile("payloads.db", nil)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Put([]byte(time.Now().String()), payload.Data, nil)
	if err != nil {
		return err
	}

	// set response data
	response.Data = payload.Data
	return nil
}

func BenchmarkProcessPayload(b *testing.B) {
	// create RPC server
	payloadService := new(PayloadService)
	rpc.Register(payloadService)
	listener, err := net.Listen("tcp", ":11111")
	if err != nil {
		panic(err)
	}
	go rpc.Accept(listener)

	// create RPC client
	client, err := rpc.Dial("tcp", "localhost:11111")
	if err != nil {
		panic(err)
	}

	// benchmark with 500 bytes payload
	iters, bytes_size := 34, 333
	payload := make([]byte, bytes_size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < iters; j++ {
			var response PayloadResponse
			err = client.Call("PayloadService.Process", Payload{Data: payload}, &response)
			if err != nil {
				panic(err)
			}
		}
	}
	b.StopTimer()

	// benchmark with 1 MB payload
	iters, bytes_size = 1000, 1e6
	payload = make([]byte, bytes_size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < iters; j++ {
			var response PayloadResponse
			err = client.Call("PayloadService.Process", Payload{Data: payload}, &response)
			if err != nil {
				panic(err)
			}
		}
	}
	b.StopTimer()

	// close client and server
	client.Close()
	listener.Close()
}

func main() {
	// create RPC server
	payloadService := new(PayloadService)
	rpc.Register(payloadService)
	listener, err := net.Listen("tcp", ":11111")
	if err != nil {
		panic(err)
	}
	go rpc.Accept(listener)

	// create RPC client
	client, err := rpc.Dial("tcp", "localhost:11111")
	if err != nil {
		panic(err)
	}

	// benchmark with 500 bytes payload
	iters, bytes_size := 34, 333
	payload := make([]byte, bytes_size)
	start := time.Now()
	for i := 0; i < iters; i++ {
		var response PayloadResponse
		err = client.Call("PayloadService.Process", Payload{Data: payload}, &response)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("Response: %v\n", response.Data)
	}
	elapsed := time.Since(start)
	fmt.Printf("%d bytes payload total %dmb took %s\n", bytes_size, iters*bytes_size/1000, elapsed)

	// benchmark with 1 MB payload
	iters, bytes_size = 1000, 1e6
	payload = make([]byte, bytes_size)
	start = time.Now()
	for i := 0; i < 10; i++ {
		var response PayloadResponse
		err = client.Call("PayloadService.Process", Payload{Data: payload}, &response)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("Response: %v\n", len(response.Data))
	}
	elapsed = time.Since(start)
	fmt.Printf("1mb payload total %dmb took %s\n", iters*bytes_size/1000000, elapsed)

	// // read and print stored payloads from LevelDB
	// db, err := leveldb.OpenFile("payloads.db", nil)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// iter := db.NewIterator(&util.Range{Start: nil, Limit: nil}, nil)
	// for iter.Next() {
	// 	fmt.Printf("Stored payload: %s\n", iter.Value())
	// }
	// iter.Release()

	time.Sleep(1 * time.Second)
	// close client and server
	client.Close()
	//listener.Close()
}
