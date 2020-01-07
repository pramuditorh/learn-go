### Variable

A variable is a **storage location**, with a specific type and an associated name. 

Variables in Go are created by first using the **var** keyword, then specifying the variable name eg: `myName` and the type eg: `string`, and finally, assigning a value to the variable. Assigning a value is optional, so we could use **two** statements, like this:

```go
var x string = "Hello, World"
```
or
```go
var x string 
x = "Hello, World"
```

### Constant

Go also has support for **constants**. Constants are essentially **variables whose values cannot be changed** later. They are created in the same way you create variables, but instead of using the **var** keyword we use the **const** keyword:

```go
const x string = "Hello, World"
```

If we force to changed, it will be error, like this:

```go
x = "Some other string"
```

Result in compile-time error:
```sh
.\main.go:7: cannot assign to x
```

### Concatenation

When we see `x = x + "second"`, we should read it as "assign the concatenation of the value of the variable x and the string literal second to the variable x." The right side of the `=` is done first and the result is then assigned to the left side of the `=`.

```go
var x string
x = "first "
fmt.Println(x)
x = x + "second"
fmt.Println(x)
```

The `x = x + y` form is so common in programming that Go has a special assignment statement: `+=`. We could have written `x = x + "second"` as `x += "second"` and it would have done the same thing (other operators can be used the same way).

```go
var x string
x = "first "
fmt.Println(x)
x += "second"
fmt.Println(x)
```

### Scope

Notice that we moved the variable outside of the main function. This means that other functions can access this variable:

```go
package main

import "fmt"

var x string = "Hello, World"

func main() {
    fmt.Println(x)
}

func f() {
    fmt.Println(x)
}
```

The `f` function now has access to the `x` variable. 

Now suppose that we wrote this instead:

```go
func main() {
    var x string = "Hello, World"
    fmt.Println(x)
}

func f() {
    fmt.Println(x)
}
```

```sh
.\main.go:11: undefined: x
```

The compiler is telling you that the `x` variable inside of the `f` function **doesn’t exist**. It only exists **inside of the main function**. The range of places where you are allowed to use `x` is called the scope of the variable. According to the language specification, **"Go is lexically scoped using blocks."** Basically, this means that the variable exists within the nearest curly braces ({ }), or block, including any nested curly braces (blocks), **but not outside of them.**

### How To Name Variable

Naming a variable properly is an important part of software development. Names **must** start with a letter and may contain **letters, numbers, or the underscore symbol (_).** The Go compiler doesn’t care what you name a variable, but you should **choose names that clearly describe the variable’s purpose.**

Like this:

```go
myAge := 22
fmt.Println("My age is", myAge)
```

Not like this:

```go
x := 22
fmt.Println("My age is", x)
```