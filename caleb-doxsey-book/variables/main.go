package main

import "fmt"

var globalScope string = "hi from outside func main"

func main() {
	var x string
	x = "first "
	fmt.Println(x) // print first
	x = x + "second "
	fmt.Println(x) // print first second
	x += "third"
	fmt.Println(x) // print first second third

	var i string = "hello"
	var j string = "world"
	fmt.Println(i == j) // print false
	j = "hello"
	fmt.Println(i == j) // print true

	// const variable declaration
	const z string = 7
	z = 10 // error. const is immutable

	// Multiple variables
	var (
		a = 1
		b = 2
		c = 3
	)

	const (
		d = 4
		e = 5
	)

	// shorter variable declaration
	k := 5
	var l = 6
	fmt.Println(k, l) // print 5 6

	// variable scope
	// local scope variable
	var localScope string = "hi from inside func main"
	fmt.Println(localScope)  // print localScope
	fmt.Println(globalScope) // print globalScope
}

func f() {
	fmt.Println(localScope) // It will print error
}
