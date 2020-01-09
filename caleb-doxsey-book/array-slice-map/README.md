- ### Arrays

    An array is a numbered sequence of elements of a single type with a fixed length. In Go, they look like this:

    ```go
    package main

    import "fmt"

    func main() {
        var x [5]int
        x[4] = 100
        fmt.Println(x)
    }
    ```
    
    The program above prints result like this:

    ```go
    [0 0 0 0 100]
    ```

    `x[4] = 100` should be read **"set the fifth element of the array x to 100."** Arrays are indexed starting from 0. Arrays are accessed in a similar way. We could change `fmt.Println(x)` to `fmt.Println(x[4])` and we would get 100.

    Simple program that use array:

    ```go
    var total float64 = 0
    for _, value := range x {
        total += value
    }
    fmt.Println(total / float64(len(x)))
    ```
    
    A single underscore (_) is **used to tell the compiler that we don’t need this.**

    Because the length of an array is part of its type name, **working with arrays can be a little cumbersome.** **Adding or removing elements** as we did here **requires also changing the length inside the brackets.** Because of this and other limitations, you **rarely see arrays used directly in Go code.** Instead, you will usually **use a slice**, which is a type built on top of an array.

- ### Slices

    A slice is a segment of an array. Like arrays, slices are **indexable** and **have a length.** Unlike arrays, this **length is allowed to change.** Here’s an example of a slice:

    ```go
    x := make([]float64, 5, 10)
    ```

    This creates a slice that is associated with an underlying `float64` array of length 5. Slices are always associated with some array, and although they can never be longer than the array, they can be smaller. The `make` function also allows a third parameter which set the max length of slice.

    - Append

    `append` adds elements onto the end of a slice. If there’s sufficient capacity in the underlying array, the element is placed after the last element and the length is incremented. However, if there is not sufficient capacity, a new array is created, all of the existing elements are copied over, the new element is added onto the end, and the new slice is returned.

    - Copy

    `copy` takes two arguments: `dst` and `src`. All of the entries in src are copied into dst overwriting whatever is there. If the lengths of the two slices are not the same, the smaller of the two will be used.


- ### Maps

    A `map` is an unordered collection of key-value pairs (maps are also sometimes called associative arrays, hash tables, or dictionaries). Maps are used to look up a value by its associated key. Here’s an example of a map in Go:

    ```go
    var x map[string]int
    x["key"] = 10
    fmt.Println(x)
    ```