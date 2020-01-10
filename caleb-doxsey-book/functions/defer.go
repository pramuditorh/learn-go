package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

/*
This has three advantages:

It keeps our Close call near our Open call so it’s easier to understand.
If our function had multiple return statements (perhaps one in an if and one in an else), Close will happen before both of them.
Deferred functions are run even if a runtime panic occurs.
*/
