package main

import (
	"strconv"
	"syscall/js"
)

// interface{} in argument to js.FuncOf
// 定義函數 js.FuncOf(add) 需 return interface{}
func add(this js.Value, i []js.Value) interface{} {
	// <button onClick="add('value1', 'value2', 'result');" id="addButton">Add</button>
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
	return int1 - int2
}

func subtract(this js.Value, i []js.Value) interface{} {
	// <button onClick="subtract('value1', 'value2', 'result');" id="subtractButton">Subtract</button>
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
	return int1 - int2
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized") // console.log

	// register functions
	registerCallbacks()
	<-c
}
