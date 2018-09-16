package main

import "fmt"

func main() {
	fmt.Println("Hello, WebAssembly!")
}

// GOOS=js GOARCH=wasm go build -o main.wasm
// cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
// goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
