package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"
)

type DataA struct {
	KeyA   string `json:"key_a"`
	ValueA string `json:"value_a"`
}

type DataB struct {
	KeyB   string `json:"key_b"`
	ValueB string `json:"value_b"`
}

type Template struct {
	DataA []DataA `json:"data_a"`
	DataB DataB   `json:"data_b"`
}

func TestServer(t *testing.T) {
	go fasthttp.ListenAndServe(":8081", func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Method()) != "POST" {
			ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			return
		}

		var template Template
		err := json.Unmarshal(ctx.PostBody(), &template)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		fmt.Printf("Received template: %+v\n", template)
	})

	// dataA := DataA{
	// 	KeyA:   "KeyA",
	// 	ValueA: "ValueA",
	// }

	dataB := DataB{
		KeyB:   "KeyB",
		ValueB: "ValueB",
	}

	template := Template{
		//DataA: []DataA{dataA},
		//DataB: []DataB{dataB},
		DataB: dataB,
	}

	data, err := json.Marshal(template)
	if err != nil {
		t.Fatalf("Failed to marshal data: %v", err)
	}

	resp, err := http.Post("http://localhost:8081", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Received non-200 response: %d", resp.StatusCode)
	}
}
