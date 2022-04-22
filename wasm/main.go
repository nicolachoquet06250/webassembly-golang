package main

import (
	"log"
	"syscall/js"
	"net/http"
	"io/ioutil"
	calculate "github.com/nicolachoquet06250/webassembly-golang/wasm/calculate"
	httprequest "github.com/nicolachoquet06250/webassembly-golang/wasm/http-request"
)

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

func MakeHttpRequest(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	println(sb)
}

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