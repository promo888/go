package main


import (
	//"fmt"
	"net"
	"net/rpc"
	"time"
	"testing"
	"github.com/syndtr/goleveldb/leveldb"
	//"github.com/syndtr/goleveldb/leveldb/util"
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
	response.Data = payload.Data  //[]byte("OK")//payload.Data
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

	// b.ResetTimer()
	// // benchmark with 1 MB payload
	// iters, bytes_size = 1000, 1e6
	// payload = make([]byte, bytes_size)
	// b.ResetTimer()
	// for i := 0; i < b.N; i++ {
	// 	for j := 0; j < iters; j++ {
	// 		var response PayloadResponse
	// 		err = client.Call("PayloadService.Process", Payload{Data: payload}, &response)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }
	// b.StopTimer()

	// close client and server
	client.Close()
	listener.Close()
}
