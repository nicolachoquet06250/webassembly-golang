package main

import (
	"strconv"
	"syscall/js"
	"net/http"
	"io/ioutil"
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

func fetch(url string, resolve js.Value, reject js.Value) any {
	// Run this code asynchronously
	go func() {
		// Make the HTTP request
		res, err := http.DefaultClient.Get(url)
		if err != nil {
			// Handle errors: reject the Promise if we have an error
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(err.Error())
			reject.Invoke(errorObject)
			return
		}
		defer res.Body.Close()

		// Read the response body
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			// Handle errors here too
			errorConstructor := js.Global().Get("Error")
			errorObject := errorConstructor.New(err.Error())
			reject.Invoke(errorObject)
			return
		}

		// "data" is a byte slice, so we need to convert it to a JS Uint8Array object
		arrayConstructor := js.Global().Get("Uint8Array")
		dataJS := arrayConstructor.New(len(data))
		js.CopyBytesToJS(dataJS, data)

		// Create a Response object and pass the data
		responseConstructor := js.Global().Get("Response")
		response := responseConstructor.New(dataJS)

		// Resolve the Promise
		resolve.Invoke(response)
	}()

	// The handler of a Promise doesn't return any value
	return nil
}

func registerCallbacks() {
    js.Global().Set("add", js.FuncOf(add))

    js.Global().Set("subtract", js.FuncOf(subtract))

	js.Global().Set("MyGoFunc", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Get the URL as argument
		// args[0] is a js.Value, so we need to get a string out of it
		requestUrl := args[0].String()

		// Handler for the Promise
		// We need to return a Promise because HTTP requests are blocking in Go
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			return fetch(requestUrl, resolve, reject)
		})

		// Create and return the Promise object
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	}))

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

	js.Global().Call("MyGoFunc", "https://jsonplaceholder.typicode.com/todos/1").
		Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return args[0].Call("json")
		})).Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			json := args[0]
			json_str := js.Global().Get("JSON").Call("stringify", json)
			println(json_str.String())
			return nil
		}))
    <-c
}