package httprequest

import (
	"io/ioutil"
	"log"
	"net/http"
	"syscall/js"
)

type Value = js.Value

func fetch(url string, resolve Value, reject Value) any {
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

func makeHttpRequest(url string) {
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

func makeHttpRequestForJS(this Value, args []Value) interface{} {
	// Get the URL as argument
	// args[0] is a js.Value, so we need to get a string out of it
	requestUrl := args[0].String()

	// Handler for the Promise
	// We need to return a Promise because HTTP requests are blocking in Go
	handler := js.FuncOf(func(this Value, args []Value) interface{} {
		resolve := args[0]
		reject := args[1]

		return fetch(requestUrl, resolve, reject)
	})

	// Create and return the Promise object
	promiseConstructor := js.Global().Get("Promise")
	thenCb := js.FuncOf(func(this Value, args []Value) interface{} { return args[0].Call("json") })
	return promiseConstructor.New(handler).Call("then", thenCb)
}

func registerMakeHttpRequestWithName(name ...string) {
	if len(name) == 0 {
		name[0] = "MakeHttpRequest"
	}

	js.Global().Set(name[0], js.FuncOf(makeHttpRequestForJS))
}

func callMakeHttpRequestFromName(url string, name ...string) {
	if len(name) == 0 {
		name[0] = "MakeHttpRequest"
	}

	requestCb := js.FuncOf(func(this Value, args []Value) interface{} {
		json := args[0]
		json_str := js.Global().Get("JSON").Call("stringify", json)

		println(json_str.String())

		return nil
	})

	js.Global().Call(name[0], url).Call("then", requestCb)
}

func RegisterCallbacks() {
	// requêtes http via js
	var funcName string = "MyGoFunc"
	registerMakeHttpRequestWithName(funcName)
	callMakeHttpRequestFromName("https://jsonplaceholder.typicode.com/todos/1", funcName)

	// requêtes http via Go
	makeHttpRequest("https://jsonplaceholder.typicode.com/posts/1")
}
