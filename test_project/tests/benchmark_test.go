// package main

// import (
// 	"encoding/json"
// 	"testing"

// 	"github.com/valyala/fasthttp"
// )

// type Data struct {
// 	Key   string `json:"key"`
// 	Value string `json:"value"`
// }

// func BenchmarkPostJSON(b *testing.B) {
// 	//b.N = 1000
// 	//reqAmount := 1000

// 	url := "https://example.com/api/endpoint"
// 	data := Data{
// 		Key:   "example_key",
// 		Value: "example_value",
// 	}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		b.Fatalf("failed to marshal data: %v", err)
// 	}

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ { //b.N reqAmount
// 		req := fasthttp.AcquireRequest()
// 		req.SetRequestURI(url)
// 		req.Header.SetMethod("POST")
// 		req.Header.SetContentType("application/json")
// 		req.SetBody(jsonData)

// 		resp := fasthttp.AcquireResponse()
// 		err = fasthttp.Do(req, resp)
// 		if err != nil {
// 			b.Fatalf("failed to make request: %v", err)
// 		}

// 		var respData Data
// 		err = json.Unmarshal(resp.Body(), &respData)
// 		if err != nil {
// 			b.Fatalf("failed to unmarshal response data: %v", err)
// 		}

// 		fasthttp.ReleaseRequest(req)
// 		fasthttp.ReleaseResponse(resp)
// 	}
// }

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"
)

type DataAA struct {
	KeyA   string `json:"key_a"`
	ValueA string `json:"value_a"`
}

type DataBB struct {
	KeyB   string `json:"key_b"`
	ValueB string `json:"value_b"`
}

type TemplateT struct {
	DataA []DataAA `json:"data_a"`
	DataB []DataBB `json:"data_b"`
}

func BenchmarkServer(b *testing.B) {
	go fasthttp.ListenAndServe(":8081", func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Method()) != "POST" {
			ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			return
		}

		var template TemplateT
		err := json.Unmarshal(ctx.PostBody(), &template)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		fmt.Printf("Received template: %+v\n", template.DataA[0].KeyA)
	})

	dataA := DataAA{
		KeyA:   "KeyA",
		ValueA: "ValueA",
	}

	dataB := DataBB{
		KeyB:   "KeyB",
		ValueB: "ValueB",
	}

	template := TemplateT{
		DataA: []DataAA{dataA},
		DataB: []DataBB{dataB},
	}

	data, err := json.Marshal(template)
	if err != nil {
		b.Fatalf("Failed to marshal data: %v", err)
	}

	b.ResetTimer()
	b.SetParallelism(10000)
	for i := 0; i < 1000; i++ {
		fmt.Print(i, "-")
		resp, err := http.Post("http://localhost:8081", "application/json", bytes.NewReader(data))
		if err != nil {
			b.Fatalf("Failed to send request: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			b.Fatalf("Received non-200 response: %d", resp.StatusCode)
		}
	}
}
