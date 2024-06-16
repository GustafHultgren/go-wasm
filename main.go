package main

import (
	"fmt"
	"os"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("Webassembly loaded")

	utils := js.ValueOf(map[string]interface{}{
		"hello": js.FuncOf(hello),
		"sum":   js.FuncOf(sum),
	})

	js.FuncOf(hello)

	js.Global().Set("__wasm__", utils)

	<-c
}

func hello(this js.Value, args []js.Value) interface{} {
	for _, input := range args {
		fmt.Printf("Hello %s\n", input.String())
	}

	return nil
}

func sum(this js.Value, inputs []js.Value) interface{} {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Error:\n%v\n", r)
		}
	}()

	total := float64(0)

	for _, input := range inputs {
		total += input.Float()
	}

	return js.ValueOf(total)
}
