const go = new Go();

const gist_url = 'https://gist.githubusercontent.com/nicolachoquet06250/2f001ee83be8a23e0ec605ec3a9de016';

const filename = 'testnum.wasm';

const wasm_file_path = window.location.host.indexOf('.github.io') !== -1
    ? `${gist_url}/raw/0cc8148a2b5afead6a003656c4cae8b19ef43ce9/${filename}` : `wasm/${filename}`;

WebAssembly.instantiateStreaming(fetch(wasm_file_path), go.importObject)
    .then(async result => {
        await go.run(result.instance);
    })
