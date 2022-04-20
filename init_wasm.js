const go = new Go();
let mod, inst;

WebAssembly.instantiateStreaming(fetch("testnum.wasm"), go.importObject)
    .then(
        async result => {
            mod = result.module;
            inst = result.instance;
            await go.run(inst);
        }
    );