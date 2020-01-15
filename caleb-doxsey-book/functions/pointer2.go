package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("Before foo(x)")
	fmt.Println("value: ", x)
	fmt.Println("address: ", &x)
	foo(&x)
	fmt.Println("After foo(x)")
	fmt.Println("value: ", x)
	fmt.Println("address: ", &x)
}

func foo(y *int) {
	fmt.Println("Inside foo() before y")
	fmt.Println("value: ", *y)
	fmt.Println("address: ", y)
	*y = 43
	fmt.Println("Inside foo() after y")
	fmt.Println("value: ", *y)
	fmt.Println("address: ", y)
}
