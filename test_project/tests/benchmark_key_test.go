package main

import (
	"crypto/rand"
	//"fmt"
	"testing"

	"golang.org/x/crypto/ed25519"
	//"fmt"
	//"time"
	//"cpu"
	//"runtime"
)

func BenchmarkVerifyKey(b *testing.B) {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)

	data := []byte("test data")
	signature := ed25519.Sign(privateKey, data)

	// fmt.Println("publicKey: ", publicKey)
	// fmt.Println("privateKey: ", privateKey)
	// fmt.Println("data: ", data)
	// fmt.Println("signature: ", signature)


	b.SetParallelism(10000)
	//fmt.Println("b.N", b.N)
	//runtime.GOMAXPROCS(1)
	//b.SetParallelism(1)
	b.ResetTimer()

	for i := 0; i < 1000; i++ { // b.N
		ed25519.Verify(publicKey, data, signature)
		//cnt=i
	}
	//fmt.Println("cnt", cnt)



	// //cpu.SetMaxProcs(1)
	// runtime.GOMAXPROCS(1)
	// done := make(chan bool)
	// go func() {
	// 	//b.SetParallelism(1)
	// 	//b.N = 1000
	// 	b.ResetTimer()
	// 	for i := 0; i < 1000; i++ {
	// 		ed25519.Verify(publicKey, data, signature)
	// 	}
	// 	done <- true
	// }()

	// select {
	// case <-done:
	// 	return
	// case <-time.After(time.Second):
	// 	b.Fatalf("Benchmark timed out")
	// }
}
