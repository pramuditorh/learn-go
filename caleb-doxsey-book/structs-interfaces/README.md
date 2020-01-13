- ### Structs

    A struct is a type that contains named fields. For example, we could represent a circle like this:

    ```go
    type Circle struct {
    x, y, r float64
    }
    ```

    The `type` keyword introduces a new type. Followed by type name (`Circle`). Keyword `struct` defining a `struct` type.

    - Initialization

        Create an instance of `Circle` type. Go has variety ways to initialize it:

        ```go
        var c Circle
        ```

        This create local `Circle` variable and the values is set by default to zero (0 for `int`, 0.0 for `float`, `""` for `string`, `nil` for pointers, etc).

        ```go
        c := new(Circle)
        ```

        This allocates memory for all fields, sets each of them to zero, and returns pointer to the struct (`*Circle`). Pointers are often to used with structs so that functions can modify the value.

        Using `new` in this way is **somewhat uncommon.** More typically, we want to give each of the fields an initial value. We can do this in two ways.

        ```go
        c := Circle{x: 0, y: 0, r: 5}
        ```

        ```go
        c := Circle{0, 0, 5}
        ```

        This creates the same `Circl`e as the previous example. If you want a pointer to the struct, use `&`:

        ```go
        c := &Circle{0, 0, 5}
        ```

    - Fields

        We can access fields by using `.` operator:

        ```go
        fmt.Println(c.x, c.y, c.r)
        c.x = 10
        c.y = 5
        ```

        Example:

        ```go
        func circleArea(c Circle) float64 {
            return math.Pi * c.r*c.r
        }

        func main() {
            c := Circle{0, 0, 5}
            fmt.Println(circleArea(c))
        }
        ```

        **One thing to remember is that arguments are always copied in Go.** If we attempted to modify one of the fields inside of the `circleArea` function, it would **not** modify the original variable. Because of this, we would typically **write the function using a pointer** to the `Circle`:

        ```go
        func circleArea(c *Circle) float64 {
            return math.Pi * c.r*c.r
        }

        func main() {
            c := Circle{0, 0, 5}
            fmt.Println(circleArea(&c))
        }
        ```

<br>

- ### Methods

    ```go
    func (c *Circle) area() float64 {
        return math.Pi * c.r*c.r
    }
    ```

    **In between the keyword func and the name of the function, we’ve added a *receiver*.** The receiver is like a parameter—it has a name and a type—but by creating the function in this way, it allows us to call the function using the `.` operator:

    ```go
    fmt.Println(c.area())
    ```

    This is much easier to read. We no longer need the `&` operator (Go automatically knows to pass a pointer to the circle for this method), **and because this function can only be used with** `Circle`, we can rename the function to just area.

    New example (with `Rectangle`):

    ```go
    type Rectangle struct {
        x1, y1, x2, y2 float64
    }

    func (r *Rectangle) area() float64 {
        l := distance(r.x1, r.y1, r.x1, r.y2)
        w := distance(r.x1, r.y1, r.x2, r.y1)
        return l * w
    }

    func main() {
        r := Rectangle{0, 0, 10, 10}
        fmt.Println(r.area())
    }
    ```

    - Embedded Types

        A struct’s fields usually represent the has-a relationship (e.g., a `Circle` has a radius). A field declared **with a type but no explicit field name is an anonymous field,** also called an **embedded field** or an embedding of the type in the struct. 

        ```go
        type Person struct {
            Name string
        }

        func (p *Person) Talk() {
            fmt.Println("Hi, my name is", p.Name)
        }
        ```

        Supposed we want to create a new struct called `Android`:

        ```go
        type Android struct {
            Person
            Model string
        }
        ```

        We use the type (`Person`) and don’t give it a name. When defined this way, the `Person` struct can be accessed using the type name:

        ```go
        a := new(Android)
        a.Person.Talk()
        ```

        Or directly call `Person` methods on `Android`:

        ```go
        a := new(Android)
        a.Talk()
        ```

<br>

- ### Interfaces

    ```go
    type Shape interface {
        area() float64
    }
    ```

    Like a struct, an interface is created using the `type` keyword, followed by a name and the keyword `interface`. But instead of defining fields, we define a **method set. A method set is a list of methods that a type must have in order to implement the interface.**

    In our case, both `Rectangle` and `Circle` have area methods that return `float64`, so both types implement the `Shape` interface.

    ```go
    func totalArea(shapes ...Shape) float64 {
        var area float64
        for _, s := range shapes {
            area += s.area()
        }
        return area
    }

    func main() {
        fmt.Println(totalArea(&c, &r))
    }
    ```

    All `totalArea` knows about each shape is that it has an `area` method.

    **Interfaces can also be used as fields.** Consider a `MultiShape` that is made up of several smaller shapes:

    ```go
    type MultiShape struct {
        shapes []Shape
    }

    multiShape := MultiShape{
        shapes: []Shape{
            Circle{0, 0, 5},
            Rectangle{0, 0, 10, 10},
        },
    }
    ```

    ```go
    func (m *MultiShape) area() float64 {
        var area float64
        for _, s := range m.shapes {
            area += s.area()
        }
        return area
    }
    ```

    Now a `MultiShape` can contain `Circle`, `Rectangle`, or even other `MultiShape`.

    Interfaces are particularly useful as software projects grow and become more complex. **They allow us to hide the incidental details of implementation** (e.g., the fields of our struct), which makes it easier to reason about software components in isolation.