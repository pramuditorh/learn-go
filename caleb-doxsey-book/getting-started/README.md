```go
package main 
```
Known as **package declaration**. Every Go programs start with it. Packages are Go's ways to organising and reusing code.

There are 2 types of Go programs:
- **executables**
- **libraries**

Executables are kind of programs that can be run directly from terminal (like .exe on Windows).
Libraries are **collections of code** that packaged together so it can be used on another programs.

<br>
<br>

```go
import "fmt"
```

`import` is **ways to include code from other package** like `main` to use it in our programs. 
`fmt` is use for formatting for input and output.

<br>
<br>

```go
func main() {
	fmt.Println("Hello World!")
}
```

Functions are building block of Go programs. All functions in Go started with `func` followed by the name of the function, which is `main` in this code. Then followed by `()` which is filled by parameters, in this code we have 0 parameter. Then followed by `{}` which is the body of the function. Name `main` is special because it is the function that always be called when run programs.

<br>
<br>

```go
fmt.Println("Hello World!")
```
This statement is made of three components. First, we access another function inside of the `fmt` package called `Println`. `Println` means **print line**. Then we create a new string that contains `Hello, World` and execute that function with the string as the first and only argument.