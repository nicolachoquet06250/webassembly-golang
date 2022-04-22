module github.com/nicolachoquet06250/webassembly-golang/wasm

go 1.18

replace github.com/nicolachoquet06250/webassembly-golang/wasm/calculate => ../wasm/calculate

replace github.com/nicolachoquet06250/webassembly-golang/wasm/http-request => ../wasm/http-request

require (
	github.com/dennwc/dom v0.3.0 // indirect
	honnef.co/go/js/dom/v2 v2.0.0-20210725211120-f030747120f2
)

require (
	github.com/nicolachoquet06250/webassembly-golang/wasm/calculate v0.0.0-00010101000000-000000000000 // indirect
	github.com/nicolachoquet06250/webassembly-golang/wasm/http-request v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
)
