const go = new Go();

const gist_url = 'https://gist.githubusercontent.com/nicolachoquet06250/2f001ee83be8a23e0ec605ec3a9de016';

const filename = 'testnum.wasm';

const wasm_file_path = window.location.host.indexOf('.github.io') !== -1
    ? `${gist_url}/raw/1f067358c02bde72a12b607598b5e5d817067a4e/${filename}` : `wasm/${filename}`;

WebAssembly.instantiateStreaming(fetch(wasm_file_path), go.importObject)
    .then(async result => {
        await go.run(result.instance);
    })
