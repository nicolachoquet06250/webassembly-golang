package main

import (
	"github.com/nicolachoquet06250/webassembly-golang/wasm/calculate"
	httprequest "github.com/nicolachoquet06250/webassembly-golang/wasm/http-request"
	"syscall/js"
)

type Value = js.Value

func registerCallbacks() {
	calculate.RegisterCallbacks()

	calculationButton := js.Global().Get("document").Call("querySelector", "#calcul")
	handleClick := js.FuncOf(func(this Value, _ []Value) interface{} {
		if calculate.GetOperator() == "add" {
			return js.Global().Get("window").Call("add", "value1", "value2", "result")
		} else if calculate.GetOperator() == "substract" {
			return js.Global().Get("window").Call("subtract", "value1", "value2", "result")
		}

		return 0
	})
	calculationButton.Call("addEventListener", "click", handleClick)

	httprequest.RegisterCallbacks()
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
