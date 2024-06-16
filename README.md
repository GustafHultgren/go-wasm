# Minimal golang webassembly example

## Resources
https://github.com/golang/go/wiki/WebAssembly
https://permify.co/post/wasm-go/
https://thenewstack.io/webassembly-and-go-a-guide-to-getting-started-part-1/

## Setup
Copy javascript support file
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

Build command
```
GOOS=js GOARCH=wasm go build -o main.wasm
```

Hot build (ignore errors)
```
air -c air.toml
```
