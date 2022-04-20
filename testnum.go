package main

import (
	"strconv"
	"syscall/js"
	"honnef.co/go/js/dom/v2"
)

func add(this js.Value, i []js.Value) any {
	value1Selector := "#" + i[0].String()
	value2Selector := "#" + i[1].String()
	resultSelector := "#" + i[2].String()

	value1 := js.Global().Get("document").Call("querySelector", value1Selector).Get("value").String()
    value2 := js.Global().Get("document").Call("querySelector", value2Selector).Get("value").String()

    int1, _ := strconv.Atoi(value1)
    int2, _ := strconv.Atoi(value2)

	resultEl := *(dom.GetWindow().Document().QuerySelector(resultSelector)).(*dom.HTMLInputElement)

    resultEl.Set("value", int1 + int2)

	println("add: ", resultEl.Value())

	return ""
}

func subtract(this js.Value, i []js.Value) any {
	value1Selector := "#" + i[0].String()
	value2Selector := "#" + i[1].String()
	resultSelector := "#" + i[2].String()

	value1 := js.Global().Get("document").Call("querySelector", value1Selector).Get("value").String()
    value2 := js.Global().Get("document").Call("querySelector", value2Selector).Get("value").String()

    int1, _ := strconv.Atoi(value1)
    int2, _ := strconv.Atoi(value2)

	resultEl := *(dom.GetWindow().Document().QuerySelector(resultSelector)).(*dom.HTMLInputElement)

    resultEl.Set("value", int1 - int2)

	println("subtract: ", resultEl.Value())

	return ""
}

func GetOperator() string {
	return (*(dom.GetWindow().Document().QuerySelector("#operator")).(*dom.HTMLSelectElement)).Value()
}

func registerCallbacks() {
    js.Global().Set("add", js.FuncOf(add))

    js.Global().Set("subtract", js.FuncOf(subtract))

	calculHandler := js.FuncOf(func (this js.Value, _ []js.Value) interface{} {
		if GetOperator() == "add" {
			return js.Global().Get("window").Call("add", "value1", "value2", "result")
		} else if GetOperator() == "substract" {
			return js.Global().Get("window").Call("subtract", "value1", "value2", "result")
		}
		return 0
	})
	
	calculButton := js.Global().Get("document").Call("querySelector", "#calcul")

	calculButton.Call("addEventListener", "click", calculHandler)
}

func main() {
    c := make(chan struct{}, 0)

    println("WASM Go Initialized")
    // register functions
    registerCallbacks()
    <-c
}