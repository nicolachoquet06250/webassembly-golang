package main

import (
	"syscall/js"
	calculate "github.com/nicolachoquet06250/webassembly-golang/wasm/calculate"
	httprequest "github.com/nicolachoquet06250/webassembly-golang/wasm/http-request"
)

func registerCallbacks() {
    calculate.RegisterCallbacks()

	calculButton := js.Global().Get("document").Call("querySelector", "#calcul")
	handleClick := js.FuncOf(func (this js.Value, _ []js.Value) interface{} {
		if calculate.GetOperator() == "add" {
			return js.Global().Get("window").Call("add", "value1", "value2", "result")
		} else if calculate.GetOperator() == "substract" {
			return js.Global().Get("window").Call("subtract", "value1", "value2", "result")
		}
		
		return 0
	})
	calculButton.Call("addEventListener", "click", handleClick)

	httprequest.RegisterCallbacks()
}

func main() {
    c := make(chan struct{}, 0)

    println("WASM Go Initialized")
    // register functions
    registerCallbacks()
    <-c
}