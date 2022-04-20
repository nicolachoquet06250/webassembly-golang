const go = new Go();

WebAssembly.instantiateStreaming(fetch("testnum.wasm"), go.importObject)
    .then(async result => {
        await go.run(result.instance);
    })
