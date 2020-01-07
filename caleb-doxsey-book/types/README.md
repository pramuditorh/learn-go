Go comes with several built-in data types.

- ### Numbers
    - Integers

        Integers—like their mathematical counterpart—are numbers without a decimal component. (…, −3, −2, −1, 0, 1, …) Unlike the base-10 decimal system we use to represent numbers, computers use a base-2 binary system.

        Go’s integer types are `uint8, uint16, uint32, uint64, int8, int16, int32, and int64`. 8, 16, 32, and 64 tell us how many bits each of the types use. uint means **unsigned integer** while int means **signed integer**. **Unsigned integers only contain positive numbers (or zero)**. In addition, there two alias types: `byte` (which is the same as uint8) and `rune` (which is the same as int32).

    - Floating-point Numbers

        Floating-point numbers are numbers that contain a decimal component (i.e., real numbers). For example, 1.234, 123.4, 0.00001234, and 12340000 are all floating-point numbers.

        Go has **two floating-point types**: `float32` and `float64` (also often referred to as single precision and double precision, respectively). It also has two additional types for representing complex numbers (numbers with imaginary parts): `complex64` and `complex128`. Generally, we should stick with float64 when working with floating-point numbers.

    Go supports the following standard arithmetic operators: <br>
        +&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Addition <br>
        -&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Subtraction <br>
        *&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Multiplication <br>
        /&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Division <br>
        %&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Remainder

- ### Strings

    As we saw in Chapter 1, **a string is a sequence of characters with a definite length used to represent text**. Go **strings are made up of individual bytes**, usually one for each character (characters from some languages, such as Chinese, are represented by more than one byte).    **String literals** can be created using **double quotes** or **backticks**. The difference between these is that double-quoted strings cannot contain newlines and they allow special escape sequences. For example, \n gets replaced with a newline and \t gets replaced with a tab character.

- ### Booleans

    &&&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;and <br>
    ||&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or <br>
    !&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;not <br>