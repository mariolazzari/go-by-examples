# Go by examples

Official [site](https://gobyexample.com/)

## Hello world

First program in Go.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

```sh 
go run hello-world.go
go build hello-world.go
```

## Values

Go has various *value types* including strings, integers, floats, booleans...
Here are a few basic examples.

```go
package main

import "fmt"

func main() {
    fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
```

## Variables

In Go, *variables* are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
The *:=* syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.

```go
package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
}
```


## Constants

Go supports *constants* of character, string, boolean, and numeric values.

```go
package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}
```

## For

*for* is Go’s *only looping construct*. 
Here are some basic types of for loops.

```go
package main

import "fmt"

func main() {
    // The most basic type, with a single condition.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // classic for
    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

    // range
    for i := range 3 {
        fmt.Println("range", i)
    }

    // for without a condition will loop repeatedly until you break out of the loop 
    // or return from the enclosing function.
    for {
        fmt.Println("loop")
        break
    }

    // continue to the next iteration 
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
```

## If

Branching with if and else in Go is straight-forward.
You don’t need parentheses around conditions 
There is no ternary if.

```go
package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// A statement can precede conditionals.
	// Any variables declared in this statement are available in the current and all subsequent branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```