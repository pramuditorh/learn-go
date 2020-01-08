- ### for Statement

    The `for` statement **allows us to repeat a list of statements** (a block) multiple times. Rewriting our previous program using a for statement looks like this:

    ```go
    package main

    import "fmt"

    func main() {
        i := 1
        for i <= 10 {
            fmt.Println(i)
            i = i + 1
        }
    }
    ```

    or like this:

    ```go
    func main() {
        for i := 1; i <= 10; i++ {
            fmt.Println(i)
        }
    }
    ```

    Now the **conditional expression also contains two other statements with semicolons between them**. **First,** we have the **variable initialization**, then we have the **condition to check each time,** and **finally,** we **increment the variable**. Adding 1 to a variable is so common that we have a special operator `++`; similarly, subtracting 1 can be done with `--`.

- ### if Statement

    An `if` statement is similar to a for statement in that it has a condition followed by a block. `if` statements also have an optional `else` part. If the **condition evaluates to true,** then the **block after the condition is run**; otherwise, either the block is skipped, or if the else block is present, that block is run.

    ```go
    if i % 2 == 0 {
        // divisible by 2
    } else if i % 3 == 0 {
        // divisible by 3
    } else if i % 4 == 0 {
        // divisible by 4
    }
    ```

    The **conditions are checked top down** and the **first one to result in true will have its associated block executed.** None of the other blocks will execute, even if their conditions also pass (so, for example, the number 8 is divisible by both 4 and 2, but the // divisible by 4 block will never execute because the // divisible by 2 block is done first).

- ### switch Statement

    ```go
    switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
    ```

    A `switch` statement **starts with the keyword** `switch` followed by an expression (in this case, i) and then a series of cases. The value of the expression is compared to the expression following each case keyword. If they are equivalent, then the statements following the `:` are executed.

    Like an `if` statement, **each case is checked top down and the first one to succeed is chosen.** A `switch` also **supports a default case that will happen if none of the cases matches the value** (similar to how the else works in an if statement).