if (!WebAssembly.instantiateStreaming) {
    // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        console.log(source);
        return await WebAssembly.instantiate(source, importObject);
    };
}