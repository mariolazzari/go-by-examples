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

