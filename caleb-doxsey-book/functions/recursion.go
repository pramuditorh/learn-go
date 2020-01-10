package main

import "fmt"

func factorial(xf uint) uint {
	if xf == 0 {
		return 1
	}
	return xf * factorial(xf-1)
}

func main() {
	x := uint(5)
	calculate := factorial(x)
	fmt.Println(calculate)
}
