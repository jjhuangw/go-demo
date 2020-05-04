package main

import (
	"fmt"
	"go-hello/controller"
	"runtime"
)

var a int = 123

type hotdog int

var b hotdog = 345

const (
	e = 123
	f = "foo"
)

const (
	g = iota
	h
)

func main() {
	fmt.Println("Hello" + controller.Morning)
	controller.Call()
	controller.AccessDB()
	foo()
	c := "test"
	c = "immutable"
	fmt.Println(c)
	a = int(b)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	// fmt.Printf("%T \n", a)
	// fmt.Printf("%v \n", a)

	// fmt.Printf("%v \n", g)
	// fmt.Printf("%T \n", h)

	sa := agent{
		person: person{
			first: "James",
			last:  "a",
			age:   34,
		},
		ltk: true,
	}
	bar(sa)

	// anonymous func
	f := func() {
		fmt.Println("Im a anonymous funciton")
	}

	f()

	func(x int) {
		fmt.Println("Im a anonymous funciton", x)
	}(34)

	b := bar2()
	fmt.Printf("%T \n", b)
	i := b()
	fmt.Println("return function value", i)

	h, g := even(mouse, "Ted", false)
	fmt.Println("callback function", h, g)

	ic := incrememntal()
	fmt.Println(ic())
	fmt.Println(ic())

	a1 := 23
	fmt.Println(&a1) // address
	// pointer used to reduce large data movement
	var b1 *int = &a1 // pointer
	fmt.Println(b1)
	fmt.Println(*&a1) // give you the value stored at an address when you have the address
}

// return a function
func bar2() func() int {
	return func() int {
		return 25
	}
}

// callback function, pass function as a parameter
func even(f func(s string, b bool) (string, bool), a string, c bool) (string, bool) {
	// call block
	{
		y := int(23)
		fmt.Println(y)
	}
	// fmt.Println(y) // un reach
	return f(a, c)
}

// return a function - closure
func incrememntal() func() int {
	var x int = 0

	// return the function, but the x is outside the function
	return func() int {
		x++
		return x
	}
}

// function, everything in Go is call by value
func mouse(s string, b bool) (i string, c bool) {
	return fmt.Sprintln(s), b
}

func foo() {
	// for loop
	// in a line use simeclone
	if x := 42; x != 2 {
		fmt.Println("sub function")
	}

	// switch
	switch {
	case (3 == 3), (5 == 5):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("Print ? ")
	}

	// array
	// x := []int{1, 2, 3, 4}
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println(x[1:])

	// slice
	// y := []int{5, 6, 7}
	// y = append(x, y...)
	// fmt.Println(y)

	// map
	m := map[string]int{
		"Jamse": 123,
		"May":   345,
	}

	if _, ok := m["Yex"]; !ok {
		fmt.Println("key is not existed")
	}
	m["Ted"] = 23
	delete(m, "Ted")
	for k, v := range m {
		fmt.Printf("%v %v \n", k, v)
	}

	// embeded struct
	sa := agent{
		person: person{
			first: "James",
			last:  "a",
			age:   34,
		},
		ltk: true,
	}

	p2 := person{
		first: "May",
		age:   12,
	}
	fmt.Println(sa.first, sa.last, sa.ltk)
	fmt.Println(p2)

	// anoymous structs
	a1 :=
		struct {
			name string
			age  int
		}{
			name: "momo",
			age:  12,
		}
	fmt.Println(a1)

	xx, yy := mouse("return function ", true)
	fmt.Println(xx, yy)
}

type person struct {
	first, last string
	age         int
}

type agent struct {
	person
	ltk bool
}

// interface
type human interface {
	speak()
}

// attach this method to type, extend interface
func (a agent) speak() {
	fmt.Println("Im a agent ", a.first)
}

func (a person) speak() {
	fmt.Println("Im a human ", a.first)
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I am a person-", h.(person).first) // assertion
	case agent:
		fmt.Println("I am a agent-", h.(agent).first)
	}
}
