package main

import (
	"syscall/js"
	calculate "github.com/nicolachoquet06250/webassembly-golang/wasm/calculate"
	httprequest "github.com/nicolachoquet06250/webassembly-golang/wasm/http-request"
)

func registerCallbacks() {
    js.Global().Set("add", js.FuncOf(calculate.Add))

    js.Global().Set("subtract", js.FuncOf(calculate.Subtract))

	js.Global().Set("MyGoFunc", js.FuncOf(httprequest.MakeHttpRequestForJS))

	calculHandler := js.FuncOf(func (this js.Value, _ []js.Value) interface{} {
		if calculate.GetOperator() == "add" {
			return js.Global().Get("window").Call("add", "value1", "value2", "result")
		} else if calculate.GetOperator() == "substract" {
			return js.Global().Get("window").Call("subtract", "value1", "value2", "result")
		}
		return 0
	})
	
	calculButton := js.Global().Get("document").Call("querySelector", "#calcul")

	calculButton.Call("addEventListener", "click", calculHandler)

	// requêtes http via js
	js.Global().Call("MyGoFunc", "https://jsonplaceholder.typicode.com/todos/1").
		Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return args[0].Call("json")
		})).Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			json := args[0]
			json_str := js.Global().Get("JSON").Call("stringify", json)
			println(json_str.String())
			return nil
		}))
	
	// requêtes http via Go
	httprequest.MakeHttpRequest("https://jsonplaceholder.typicode.com/posts/1")
}

func main() {
    c := make(chan struct{}, 0)

    println("WASM Go Initialized")
    // register functions
    registerCallbacks()
    <-c
}