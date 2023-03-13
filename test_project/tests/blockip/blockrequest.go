package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

const windowSize = 1 * time.Minute
const maxRequestsPerMinute = 100

type ipRequestCounter struct {
    mutex  sync.Mutex
    counts map[string][]time.Time
}

func (c *ipRequestCounter) increment(ip string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    // Remove old timestamps from the count list
    for i := 0; i < len(c.counts[ip]); i++ {
        if time.Since(c.counts[ip][i]) > windowSize {
            c.counts[ip] = c.counts[ip][i+1:]
            i--
        } else {
            break
        }
    }

    // Append the current timestamp to the count list
    c.counts[ip] = append(c.counts[ip], time.Now())
}

func (c *ipRequestCounter) count(ip string) int {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    // Remove old timestamps from the count list
    for i := 0; i < len(c.counts[ip]); i++ {
        if time.Since(c.counts[ip][i]) > windowSize {
            c.counts[ip] = c.counts[ip][i+1:]
            i--
        } else {
            break
        }
    }

    // Return the number of timestamps in the count list
    return len(c.counts[ip])
}

func (c *ipRequestCounter) reset(ip string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.counts[ip] = nil
}

// func startServer() {
// 	// Start the HTTP server
//     server := &http.Server{
//         Addr:    ":11111",
//         Handler: http.HandlerFunc(handler),
//     }
//     fmt.Println("Starting HTTP server...")
//     err := server.ListenAndServe()
//     if err != nil {
//         fmt.Println("Error starting HTTP server:", err)
//     }
// }






func main() {
    // Create the request counter
    // counter := &ipRequestCounter{
    //     counts: make(map[string][]time.Time),
    // }

    // Create the HTTP handler function
    // handler := func(w http.ResponseWriter, r *http.Request) {
    //     // Get the IP address of the client
    //     ip := r.RemoteAddr

    //     // Increment the request counter for this IP address
    //     counter.increment(ip)

    //     // Check if the IP address has exceeded the request limit
    //     if counter.count(ip) > maxRequestsPerMinute {
    //         fmt.Println("Blocking IP:", ip)

    //         // Return a 429 Too Many Requests response
    //         http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)

    //         // Reset the request count for this IP address
    //         counter.reset(ip)
    //         return
    //     }

    //     // Return a 200 OK response
    //     w.WriteHeader(http.StatusOK)
    //     w.Write([]byte("OK"))
    // }

    // // Start the HTTP server
    // server := &http.Server{
    //     Addr:    ":11111",
    //     Handler: http.HandlerFunc(handler),
    // }
    // fmt.Println("Starting HTTP server...")
    // err := server.ListenAndServe()
    // if err != nil {
    //     fmt.Println("Error starting HTTP server:", err)
    // }
	
	//go startServer()
	//time.Sleep(time.Second)


	counter := &ipRequestCounter{}
	// Create the HTTP handler function
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the IP address of the client
        ip := r.RemoteAddr

        // Increment the request counter for this IP address
        counter.increment(ip)

        // Check if the IP address has exceeded the request limit
        if counter.count(ip) > maxRequestsPerMinute {
           fmt.Println("Blocking IP:", ip)

            // Return a 429 Too Many Requests response
            http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)

            // Reset the request count for this IP address
            counter.reset(ip)
            return
        }

        // Return a 200 OK response
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })


	   // Create a test server with the handler function
	   server := httptest.NewServer(handler)
	   defer server.Close()
   
	   // Create an HTTP client for testing
	   client := &http.Client{}

	// Send 100 requests to the server from the same IP address within a minute
    ip := "192.168.1.100"
    for i := 0; i < 100; i++ {
        req, err := http.NewRequest("GET", server.URL, nil)
        if err != nil {
           log.Fatal("Error creating HTTP request:", err)
        }
        req.RemoteAddr = ip + ":12345"
        _, err = client.Do(req)
        if err != nil {
           log.Fatal("Error sending HTTP request:", err)
        }
    }

    // Send one more request to the server from the same IP address
    req, err := http.NewRequest("GET", server.URL, nil)
    if err != nil {
        log.Fatal("Error creating HTTP request:", err)
    }
    req.RemoteAddr = ip + ":12345"
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Error sending HTTP request:", err)
    }

    // Check if the server returns StatusTooManyRequests
    if resp.StatusCode != http.StatusTooManyRequests {
        log.Fatal("Expected HTTP status", http.StatusTooManyRequests, "but got", resp.StatusCode)
    }
}
