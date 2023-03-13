package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	//"net/http/httptest"
	"sync"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	//"github.com/gofiber/fiber/v2"
	//"github.com/valyala/fasthttp"
)

var blockedIPs = map[string]int{}
var mutex = &sync.Mutex{}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		mutex.Lock()
		if blockedIPs[clientIP] > 100 {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			mutex.Unlock()
			return
		}
		blockedIPs[clientIP]++
		mutex.Unlock()
		next.ServeHTTP(w, r)
	})
}

func resetIPs() {
	for {
		time.Sleep(time.Minute)
		mutex.Lock()
		blockedIPs = map[string]int{}
		mutex.Unlock()
	}
}

func startServer() {
	resetIPs() // reset blocked IPs every minute
	// log.Fatal(http.ListenAndServe(":55555", middleware()))

}

func handleGetRequest(ctx *fiber.Ctx) error {
	resetIPs() // reset blocked IPs every minute
	clientIP := ctx.Context().LocalAddr().String()
	//mutex.Lock()
	//defer mutex.Unlock()
	if blockedIPs[clientIP] > 100 {
		//mutex.Unlock()
		return ctx.Status(fiber.StatusTooManyRequests).SendString("Too Many Requests")
	}
	blockedIPs[clientIP]++
	// //mutex.Unlock()

	return ctx.Status(fiber.StatusOK).SendString("OK")
}

func startFiber() {
	app := fiber.New()
	app.Get("/", handleGetRequest)
	log.Fatal(app.Listen(":55555"))
	time.Sleep(time.Second)
}

func TestBlockedIPs(t *testing.T) {

	// 	// Send 100 requests
	// 	for i := 0; i < 100; i++ {
	// 		handler.ServeHTTP(rr, req)
	// 		if status := rr.Code; status != http.StatusOK {
	// 			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	// 		}
	// 	}

	// 	// Send one more request, it should return "Too Many Requests"
	// 	handler.ServeHTTP(rr, req)
	// 	if status := rr.Code; status != http.StatusTooManyRequests {
	// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusTooManyRequests)
	// 	}

	//go startServer()
	go fasthttp.ListenAndServe(":55555", func(ctx *fasthttp.RequestCtx) {
		//resetIPs() // reset blocked IPs every minute
		clientIP := ctx.RemoteIP().String()
		mutex.Lock()
		if blockedIPs[clientIP] > 100 {
			ctx.Error("Too Many Requests", fasthttp.StatusTooManyRequests)
			mutex.Unlock()
			return
		}
		blockedIPs[clientIP]++
		mutex.Unlock()
		fmt.Printf("Request #: %d\n", blockedIPs[clientIP])

		ctx.SetBodyString("Hello, World!")
		ctx.SetStatusCode(fasthttp.StatusOK)

	})

	//go startFiber()

	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 101; i++ {
		resp, err := http.Get("http://localhost:55555")
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Received unexpected status code: %d", resp.StatusCode)
		}
	}

	resp, err := http.Get("http://localhost:55555")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("%d - Expected to receive 429 - http.StatusTooManyRequests status", resp.StatusCode)
	}

}
