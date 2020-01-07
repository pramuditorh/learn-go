package main

import "fmt"

func main() {
	// Numbers
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1.5 + 1 = ", 1.5+1.0)
	fmt.Println("32.132 x 43.434 = ", 32132*43434)

	// Strings
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1]) // Prints binary ASCII of 'e'. This is because character 'e' is represented by binary 101
	fmt.Println("Hello, " + "World")

	// Booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
