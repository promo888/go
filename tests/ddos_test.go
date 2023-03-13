package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	WindowSize = 100
	Threshold  = 0.5
)

type Transaction2 struct {
	From        string
	To          string
	Asset       string
	Amount      float64
	Fee         float64
	Description string
	Timestamp   int64
	BlockNumber uint64
	BlockHash   string
	Transaction string
}

// detectDDoS uses a simple machine learning algorithm to detect DDoS attacks
func detectDDoS(window []*Transaction2) bool {
	// Calculate the mean and standard deviation of the transaction counts in the window
	sum := 0.0
	for _, value := range window {
		if value != nil {
			sum += 1.0
		}
	}
	mean := sum / float64(len(window))

	sum = 0.0
	for _, value := range window {
		if value != nil {
			sum += (1.0 - mean) * (1.0 - mean)
		} else {
			sum += mean * mean
		}
	}
	stdDev := sum / float64(len(window))
	stdDev = math.Sqrt(stdDev)

	// If the standard deviation is below the threshold, it's likely a DDoS attack
	if stdDev < Threshold {
		return true
	}

	// Otherwise, it's probably not a DDoS attack
	return false
}

func TestDDoSDetection(t *testing.T) {
	// Initialize a window of traffic
	window := make([]*Transaction2, WindowSize)

	// Create a test HTTP server that receives incoming transactions
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the transaction object from the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		var transaction Transaction2
		err = json.Unmarshal(body, &transaction)
		if err != nil {
			http.Error(w, "Error decoding transaction object", http.StatusBadRequest)
			return
		}

		// Update the traffic window
		window = append(window[1:], &transaction)

		// Analyze the traffic pattern to detect a DDoS attack
		if detectDDoS(window) {
			http.Error(w, "DDoS attack detected!", http.StatusTooManyRequests)
			return
		}

		// Return a success response
		w.WriteHeader(http.StatusOK)
	}))

	defer ts.Close()

	// Generate some traffic
	for i := 0; i < 1000; i++ {
		// Generate a random transaction object
		transaction := Transaction2{
			From:        "testFrom",
			To:          "testTo",
			Asset:       "testAsset",
			Amount:      123.456,
			Fee:         0.001,
			Description: "testDescription",
		}

		// Encode the transaction object as JSON
		body, err := json.Marshal(transaction)
		if err != nil {
			t.Fatalf("Error encoding transaction object: %v", err)
		}

		// Send the transaction to the test server
		resp, err := http.Post(ts.URL+"/transaction", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Error sending transaction: %v", err)
		}

		// Check if a DDoS attack was detected
		if resp.StatusCode == http.StatusTooManyRequests {
			t.Errorf("DDoS attack detected on request %d", i+1)
		}
	}
}
