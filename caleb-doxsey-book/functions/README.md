### Functions

A function (also known as a procedure or a subroutine) is an **independent section of code** that maps zero or more input parameters to zero or more output parameters.

- Return types can have names. We can name the return type like this:

    ```go
    func f2() (r int) {
        r = 1
        return
    }
    ```

- Multiple values can be returned. Go is also capable of returning multiple values from a function. Here is an example function that returns two integers:

    ```go
    func f() (int, int) {
        return 5, 6
    }

    func main() {
        x, y := f()
    }
    ```

    **Multiple values are often used to return an error value** along with the result (`x, err := f()`), **or a boolean to indicate success** (`x, ok := f()`).

<br>

### Variadic Functions

By using an ellipsis (`...`) before the type name of the last parameter, you can indicate that it takes zero or more of those parameters. 

```go
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func main() {
    fmt.Println(add(1,2,3))
}
```

<br>

### Closures

It is possible to create functions inside of functions.

```go
func main() {
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1,1))
}
```

**One way to use closure is by writing a function that returns another function,** which when called, can generate a sequence of numbers. For example, here’s how we might generate all the even numbers:

```go
func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}
func main() {
    nextEven := makeEvenGenerator()
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4
}
```

`makeEvenGenerator` **returns a function that generates even numbers.** Each time it’s called, it adds 2 to the local i variable, which—unlike normal local variables—persists between calls.

<br>

### Recursion

Recursion is a **function is able to call itself.**

```go
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
```

**Closure and recursion** are powerful programming techniques that **form the basis of a paradigm known as functional programming.**

<br>

### defer, panic, and recover

Go has a special statement called `defer` that **schedules a function call to be run after the function completes.**

`defer` is **often used when resources need to be freed** in some way. For example, when we open a file, we need to make sure to close it later. With `defer`:

```go
f, _ := os.Open(filename)
defer f.Close()
```

- ### panic and recover

    `panic` function to cause a runtime error. We can handle a runtime panic with the built-in `recover` function. `recover` stops the panic and returns the value that was passed to the call to `panic`.

    ```go
    package main

    import "fmt"

    func main() {
        defer func() {
            str := recover()
            fmt.Println(str)
        }()
        panic("PANIC")
    }
    ```

    A `panic` generally indicates a programmer error or an exceptional condition that there’s no easy way to recover from.

<br>

### Pointers

Pointers reference a location in memory where a value is stored rather than the value itself. By using a pointer (`*int`), the zero function is able to modify the original variable.

- ### * and & operators

```go
func zero(xPtr *int) {
    *xPtr = 0
}
func main() {
    x := 5
    zero(&x)
    fmt.Println(x) // x is 0
}
```

In Go, a pointer is represented using an asterisk (`*`) followed by the type of the stored value. In the zero function, `xPtr` is a pointer to an int.

An asterisk is also used to dereference pointer variables. Dereferencing a pointer gives us access to the value the pointer points to.

We use the `&` operator to find the address of a variable. `&x` returns a `*int` (pointer to an int) because `x` is an `int`. This is what allows us to modify the original variable. 