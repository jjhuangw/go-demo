package main

import (
	"fmt"
	api "go-hello/controller"
)

func main() {
	fmt.Println("Hello" + api.Morning)
	api.Call()
	foo()
}

func foo() {
	fmt.Println("sub")
}
