package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://example.com", http.StatusMovedPermanently)
}

func BenchmarkRedirect(b *testing.B) {
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		handler(rr, req)
	}
}

// func main() {
// 	fmt.Println("Starting benchmark test...")
// 	fmt.Println("Please wait...")
// 	testing.Benchmark(BenchmarkRedirect)
// 	fmt.Println("Benchmark test completed.")
// }

func BenchmarkRedirects(b *testing.B) {
	// Create a new HTTP server with the redirect handler
	srv := httptest.NewServer(http.HandlerFunc(handler))
	fmt.Println("Server started on", srv.URL)
	time.Sleep(100 * time.Millisecond)

	b.ResetTimer()
	// Benchmark the redirect handler
	fmt.Println("Starting benchmark test...")
	fmt.Println("Please wait...")
	testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res, err := http.Get(srv.URL)
			if err != nil {
				b.Fatal(err)
			}
			res.Body.Close()
		}
	})
	fmt.Println("Benchmark test completed.")
}

func dummyGet(url string) {
	return
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\nResponse status: %s from %s: ", res.Status, res.Request.URL.String())
	res.Body.Close()
}

func TestRedirects(t *testing.T) {
	// Create a new HTTP server with the redirect handler
	srv := httptest.NewServer(http.HandlerFunc(handler))
	start := time.Now()
	fmt.Println("Server started on", srv.URL)
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Starting benchmark test...")
	fmt.Println("Please wait...")
	// for i := 0; i < 10; i++ {
	// 	res, err := http.Get(srv.URL)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("\nResponse status: %s from %s: ", res.Status, res.Request.URL.String())
	// 	res.Body.Close()
	// }
	var wg sync.WaitGroup
	concurrents := 100000 //1000
	wg.Add(concurrents)   //(1000)
	for i := 0; i < concurrents; i++ {
		go func() {
			defer wg.Done()
			// res, err := http.Get(srv.URL)
			// if err != nil {
			// 	fmt.Println("Error:", err)
			// 	return
			// }
			// fmt.Printf("\nResponse status: %s from %s: ", res.Status, res.Request.URL.String())
			// res.Body.Close()
			dummyGet(srv.URL)
		}()
	}
	wg.Wait()

	fmt.Printf("\nBenchmark test completed, took: %s", time.Since(start))
}
