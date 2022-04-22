package httprequest

import (
	"log"
	"syscall/js"
	"net/http"
	"io/ioutil"
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

func MakeHttpRequestForJS(this js.Value, args []js.Value) interface{} {
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
}