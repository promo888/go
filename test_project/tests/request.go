package main

// type DataA struct {
// 	KeyA   string `json:"key_a"`
// 	ValueA string `json:"value_a"`
// }

// type DataB struct {
// 	KeyB   string `json:"key_b"`
// 	ValueB string `json:"value_b"`
// }

// type Template struct {
// 	DataA []DataA `json:"data_a"`
// 	DataB []DataB `json:"data_b"`
// }

//func main() {
// fasthttp.ListenAndServe(":8081", func(ctx *fasthttp.RequestCtx) {
// 	if string(ctx.Method()) != "POST" {
// 		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
// 		return
// 	}

// 	var template Template
// 	err := json.Unmarshal(ctx.PostBody(), &template)
// 	if err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println("Received template: %+v\n", template)
// })
//}
