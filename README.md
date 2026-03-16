# [Go by examples](https://gobyexample.com/)

[Go](https://go.dev/) is an open source programming language designed for building scalable, secure and reliable software. Please read the [official documentation](https://go.dev/doc/tutorial/getting-started) to learn more.

_Go by Example_ is a hands-on introduction to Go using annotated example programs. Check out the [first example](https://gobyexample.com/hello-world) or browse the full list below.

Unless stated otherwise, examples here assume the [latest major release Go](https://go.dev/doc/devel/release) and may use new language features. Try to upgrade to the latest version if something isn't working.

## Hello world

Our first program will print the classic “hello world” message. Here’s the full source code.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

To run the program, put the code in hello-world.go and use go run.

```sh
go run hello-world.go
```

Sometimes we’ll want to build our programs into binaries. We can do this using go build.

```sh
go build hello-world.go
```

We can then execute the built binary directly.

```sh
./hello-world
```

## Values

Go has various _value types_ including strings, integers, floats, booleans...
Here are a few basic examples.

```go
func main() {
	// Strings, which can be added together with +.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators as you’d expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
```

```sh
$ go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```

## Variables

In Go, _variables_ are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
The _:=_ syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.

```go
package main

import "fmt"

func main() {
	// var declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	// For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable,
	// e.g. for var f string = "apple" in this case.
	// This syntax is only available inside functions.
	f := "apple"
	fmt.Println(f)
}
```

```sh
$ go run variables.go
initial
1 2
true
0
apple
```

## Constants

Go supports _constants_ of character, string, boolean, and numeric values.

```go
package main

import (
    "fmt"
    "math"
)

// const declares a constant value.
const s string = "constant"

func main() {
	fmt.Println(s)

	// A const statement can also appear inside a function body.
	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call.
	// For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}
```

```sh

$ go run constant.go
constant
6e+11
600000000000
-0.28470407323754404
```

## For

_for_ is Go’s _only looping construct_.
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

	// A classic initial/condition/after for loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Another way of accomplishing the basic “do this N times” iteration is range over an integer.
	for i := range 3 {
		fmt.Println("range", i)
	}

	//for without a condition will loop repeatedly until you break out of the loop
	// or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also continue to the next iteration of the loop.
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

- You don’t need parentheses around conditions
- There is no ternary if.

```go
package main

import "fmt"

func main() {
	// Here’s a basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can have an if statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Logical operators like && and || are often useful in conditions.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// A statement can precede conditionals;
	// any variables declared in this statement are available in the current and all subsequent branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```

```sh

$ go run for.go
1
2
3
0
1
2
range 0
range 1
range 2
loop
1
3
5
```

## Switch

Switch statements express conditionals across many branches.

```go
package main

import (
    "fmt"
    "time"
)

func main() {

    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
```

```sh

$ go run switch.go
Write 2 as two
It's a weekday
It's after noon
I'm a bool
I'm an int
Don't know type string
```

## Arrays

In Go, an array is a numbered sequence of elements of a specific length.
In typical Go code, _slices_ are much more common; arrays are useful in some special scenarios.

```go
package main

import "fmt"

func main() {

// Here we create an array a that will hold exactly 5 ints.
// The type of elements and length are both part of the array’s type.
// By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
```

```sh
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
dcl: [1 2 3 4 5]
idx: [100 0 0 400 500]
2d:  [[0 1 2] [1 2 3]]
2d:  [[1 2 3] [1 2 3]]
```

## Slices

_Slices_ are an important data type in Go, giving a more powerful interface to sequences than arrays.
Check out [this great blog post](https://go.dev/blog/slices-intro) by the Go team for more details on the design and implementation of slices in Go.

```go
package main

import (
	"fmt"
	"slices"
)

func main() {
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create a slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3 (initially zero-valued).
	// By default a new slice’s capacity is equal to its length;
	// if we know the slice is going to grow ahead of time,
	// it’s possible to pass a capacity explicitly as an additional parameter to make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices support several more that make them richer than arrays.
	// One is the builtin append, which returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a “slice” operator with the syntax slice[low:high].
	// For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// The slices package contains a number of useful utility functions for slices.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
```

```sh
go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]
```

## Maps

Maps are Go’s built-in [associative data type](https://en.wikipedia.org/wiki/Associative_array) (hashes / dictionaries).
If the key doesn’t exist, the [zero value](https://go.dev/ref/spec#The_zero_value)

```go
package main

import (
	"fmt"
	"maps"
)

func main() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn’t exist, the zero value of the value type is returned.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(m)
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	// Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// The maps package contains a number of useful utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
```

```sh
go run maps.go
map: map[k1:7 k2:13]
v1: 7
v3: 0
len: 2
map: map[k1:7]
map: map[]
prs: false
map: map[bar:2 foo:1]
n == n2
```

## Functions

_Functions_ are central in Go. We’ll learn about functions with a few different examples.

```go
package main

import "fmt"

// Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	// Go requires explicit returns, i.e.
	// it won’t automatically return the value of the last expression.
	return a + b
}

// When you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// Call a function just as you’d expect, with name(args).
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
```

```sh
go run functions.go
1+2 = 3
1+2+3 = 6
```

## Multiple Return Values

Go has built-in support for _multiple return values_.
This feature is used often in idiomatic Go, for example to return both result and error.

```go
package main

import "fmt"

// The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Here we use the 2 different return values from the call with multiple assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values, use the blank identifier _.
	_, c := vals()
	fmt.Println(c)
}
```

```sh
go run multiple-return-values.go
3
7
7
```

## Variadic functions

[Variadic functions](https://en.wikipedia.org/wiki/Variadic_function) can be called with any number of trailing arguments.
For example, fmt.Println is a common variadic function.

```go
package main

import "fmt"

// Here’s a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// Within the function, the type of nums is equivalent to []int.
	// We can call len(nums), iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
```

```sh
go run variadic-functions.go
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10
```

## Closures

Go supports [anonymous functions](https://en.wikipedia.org/wiki/Anonymous_function), which can form [closures](<https://en.wikipedia.org/wiki/Closure_(computer_programming)>).
Anonymous functions are useful when you want to define a function inline without having to name it.

```go
package main

import "fmt"

// This function intSeq returns another function, which we define anonymously in the body of intSeq.
// The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value, which will be updated each time we call nextInt.
	nextInt := intSeq()

	// See the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
}
```

```sh
go run closures.go
1
2
3
1
```

## Recursion

Go supports [recursive functions](<https://en.wikipedia.org/wiki/Recursion_(computer_science)>). Here’s a classic example.

```go
package main

import "fmt"

// This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Anonymous functions can also be recursive,
	// but this requires explicitly declaring a variable with var to store the function before it’s defined.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Since fib was previously declared in main, Go knows which function to call with fib here.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
```

```sh
go run recursion.go
5040
13
```

## Range over Built-in Types

_range_ iterates over elements in a variety of built-in data structures.
Let’s see how to use range with some of the data structures we’ve already learned.

```go
package main

import "fmt"

func main() {
	// Here we use range to sum the numbers in a slice.
	// Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both the index and value for each entry.
	// Above we didn’t need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	// See Strings and Runes for more details.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
```

```sh
go run range-over-built-in-types.go
sum: 9
index: 1
a -> apple
b -> banana
key: a
key: b
0 103
1 111
```

## Pointers

Go supports [pointers](<https://en.wikipedia.org/wiki/Pointer_(computer_programming)>), allowing you to pass references to values and records within your program.

```go
package main

import "fmt"

// We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr.
// zeroval has an int parameter, so arguments will be passed to it by value.
// zeroval will get a copy of ival distinct from the one in the calling function.
func zeroval(ival int) {
	ival = 0
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The &i syntax gives the memory address of i, i.e. a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	fmt.Println("pointer:", &i)
}
```

_zeroval_ doesn’t change the i in main, but _zeroptr_ does because it has a reference to the memory address for that variable.

```sh
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
```

## Strings and runes

A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”.
In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.
[This Go blog post](https://go.dev/blog/strings) is a good introduction to the topic.

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// s is a string assigned a literal value representing the word “hello” in the Thai language.
	// Go string literals are UTF-8 encoded text.
	const s = "สวัสดี"

	// Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
	fmt.Println("Len:", len(s))

	// Indexing into a string produces the raw byte values at each index.
	// This loop generates the hex values of all the bytes that constitute the code points in s.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// To count how many runes are in a string, we can use the utf8 package.
	// Note that the run-time of RuneCountInString depends on the size of the string,
	// because it has to decode each UTF-8 rune sequentially.
	// Some Thai characters are represented by UTF-8 code points that can span multiple bytes,
	// so the result of this count may be surprising.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// This demonstrates passing a rune value to a function.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// Values enclosed in single quotes are rune literals.
	// We can compare a rune value to a rune literal directly.
	switch r {
	case 't':
		fmt.Println("found tee")
	case 'ส':
		fmt.Println("found so sua")
	}
}
```

```sh
go run strings-and-runes.go
Len: 18
e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5
Rune count: 6
U+0E2A 'ส' starts at 0
U+0E27 'ว' starts at 3
U+0E31 'ั' starts at 6
U+0E2A 'ส' starts at 9
U+0E14 'ด' starts at 12
U+0E35 'ี' starts at 15
Using DecodeRuneInString
U+0E2A 'ส' starts at 0
found so sua
U+0E27 'ว' starts at 3
U+0E31 'ั' starts at 6
U+0E2A 'ส' starts at 9
found so sua
U+0E14 'ด' starts at 12
U+0E35 'ี' starts at 15
```

## Structs

Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

```go
package main

import "fmt"

// This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the given name.
func newPerson(name string) *person {
	// Go is a garbage collected language; you can safely return a pointer to a local variable
	// it will only be cleaned up by the garbage collector when there are no active references to it.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})
	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})
	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})
	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})
	// It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))
	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers
	// the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age, s.age)

	// If a struct type is only used for a single value, we don’t have to give it a name.
	// The value can have an anonymous struct type.
	// This technique is commonly used for table-driven tests.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
```

```sh
go run structs.go
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
{Rex true}
```

## Methods

Go supports methods defined on struct types.

```go
package main

import "fmt"

type rect struct {
	width, height int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls or
	// to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
```

```sh
go run methods.go
area:  50
perim: 30
area:  50
perim: 30
```

## Interfaces

_Interfaces_ are named collections of method signatures.

```go
package main

import (
	"fmt"
	"math"
)

// Here’s a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we’ll implement this interface on rect and circle types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to implement all the methods in the interface.
// Here we implement geometry on rects.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for circles.
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call methods that are in the named interface.
// Here’s a generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Sometimes it’s useful to know the runtime type of an interface value.
// One option is using a type assertion as shown here; another is a type switch.
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The circle and rect struct types both implement the geometry interface
	// so we can use instances of these structs as arguments to measure.
	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
```

```sh
go run interfaces.go
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
circle with radius 5
```

## Enums

_Enumerated types_ (enums) are a special case of [sum types](https://en.wikipedia.org/wiki/Algebraic_data_type). An enum is a type that has a fixed number of possible values, each with a distinct name. Go doesn’t have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.

```go
package main

import "fmt"

// Our enum type ServerState has an underlying int type.
type ServerState int

// The possible values for ServerState are defined as constants.
// The special keyword iota generates successive constant values automatically;
// in this case 0, 1, 2 and so on.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// By implementing the fmt.Stringer interface, values of ServerState can be printed out or converted to strings.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// This can get cumbersome if there are many possible values.
// In such cases the stringer tool can be used in conjunction with go:generate to automate the process.
// See this post for a longer explanation.
func (ss ServerState) String() string {
	return stateName[ss]
}

// If we have a value of type int, we cannot pass it to transition
// the compiler will complain about type mismatch.
// This provides some degree of compile-time type safety for enums.
func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition emulates a state transition for a server;
// it takes the existing state and returns a new state.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Suppose we check some predicates here to determine the next state…
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
```

## Struct Embedding

Go supports embedding of structs and interfaces to express a more seamless composition of types.
This is not to be confused with _//go:embed_ which is a go directive introduced in Go version 1.16+ to embed files and folders into the application binary.

```go
package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base. An embedding looks like a field without a name
type container struct {
	base
	str string
}

func main() {

	// When creating structs with literals, we have to initialize the embedding explicitly;
	// here the embedded type serves as the field name.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// We can access the base’s fields directly on co, e.g. co.num.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	// Alternatively, we can spell out the full path using the embedded type name.
	fmt.Println("also num:", co.base.num)
	// Since container embeds base, the methods of base also become methods of a container.
	// Here we invoke a method that was embedded from base directly on co.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Embedding structs with methods may be used to bestow interface implementations onto other structs.
	// Here we see that a container now implements the describer interface because it embeds base.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
```

```sh
go run struct-embedding.go
co={num: 1, str: some name}
also num: 1
describe: base with num=1
describer: base with num=1
```

## Generics

Starting with version 1.18, Go has added support for _generics_, also known as _type parameters_.
For a more thorough explanation of this type signature, see [this blog post](https://go.dev/blog/deconstructing-type-parameters).
Note that this function exists in the standard library as [slices.Index](https://pkg.go.dev/slices#Index).

```go
package main

import "fmt"

// As an example of a generic function, SlicesIndex takes a slice of any comparable type and an element of that type
// and returns the index of the first occurrence of v in s, or -1 if not present.
// The comparable constraint means that we can compare values of this type with the == and != operators.
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// As an example of a generic type, List is a singly-linked list with values of any type.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// We can define methods on generic types just like we do on regular types,
// but we have to keep the type parameters in place. The type is List[T], not List.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements returns all the List elements as a slice.
// In the next example we’ll see a more idiomatic way of iterating
// over all elements of custom types.
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// When invoking generic functions, we can often rely on type inference.
	// Note that we don’t have to specify the types for S and E when calling SlicesIndex
	// the compiler infers them automatically.
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// though we could also specify them explicitly.
	// _ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
```

```sh
go run generics.go
index of zoo: 2
list: [10 13 23]
```

## Range over Iterators

Starting with version 1.23, Go has added support for [iterators](https://go.dev/blog/range-functions), which lets us range over pretty much anything!

```go
package main

import (
	"fmt"
	"iter"
	"slices"
)

// Let’s look at the List type from the previous example again.
// In that example we had an AllElements method that returned a slice of all elements in the list.
// With Go iterators, we can do it better - as shown below.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All returns an iterator, which in Go is a function with a special signature.
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// The iterator function takes another function as a parameter,
		// called yield by convention (but the name can be arbitrary).
		// It will call yield for every element we want to iterate over,
		// and note yield’s return value for a potential early termination.
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Iteration doesn’t require an underlying data structure, and doesn’t even have to be finite!
// Here’s a function returning an iterator over Fibonacci numbers: it keeps running as long as yield keeps returning true.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Since List.All returns an iterator, we can use it in a regular range loop.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Packages like slices have a number of useful functions to work with iterators.
	// For example, Collect takes any iterator and collects all its values into a slice.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// Standard library packages now expose iterator helpers too.
	// For example, strings.SplitSeq iterates over parts of a byte slice without first building a result slice.
	for n := range genFib() {
		// Once the loop hits break or an early return, the yield function passed to the iterator will return false.
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
```

```sh
go run range-over-iterators.go
10
13
23
all: [10 13 23]
part: go
part: by
part: example
0
1
1
2
3
5
8
```

## Errors

In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java, Python and Ruby and the overloaded single result / error value sometimes used in C. Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for other, non-error tasks.

See the documentation of the [errors package](https://pkg.go.dev/errors) and [this blog post](https://go.dev/blog/go1.13-errors) for additional details.

```go
package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have type error, a built-in interface.
func f(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("can't work with 42")
	}
	// A nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	switch arg {
	case 2:
		return ErrOutOfTea
	case 4:

		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

// We can wrap errors with higher-level errors to add context.
// The simplest way to do this is with the %w verb in fmt.Errorf.
// Wrapped errors create a logical chain (A wraps B, which wraps C, etc.)
// that can be queried with functions like errors.Is and errors.AsType.

func main() {
	for _, i := range []int{7, 42} {
		// It’s idiomatic to use an inline error check in the if line.
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			// errors.Is checks that a given error (or any error in its chain) matches a specific error value.
			// This is especially useful with wrapped or nested errors, allowing you to identify specific error types
			// or sentinel errors in a chain of errors.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
```

```sh
go run errors.go
f worked: 10
f failed: can't work with 42
Tea is ready!
Tea is ready!
We should buy new tea!
Tea is ready!
Now it is dark.
```

## Custom Errors

It’s possible to define custom error types by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.

```go
package main

import (
	"errors"
	"fmt"
)

// A custom error type usually has the suffix “Error”.
type argError struct {
	arg     int
	message string
}

// Adding this Error method makes argError implement the error interface.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		// Return our custom error.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// errors.AsType is a more advanced version of errors.Is.
	// It checks that a given error (or any error in its chain) matches a specific error type
	// and converts to a value of that type, also returning true.
	// If there’s no match, the second return value is false.
	_, err := f(42)
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
```

```sh
go run custom-errors.go
42
can't work with it
```

## Goroutines

A goroutine is a lightweight thread of execution.

```go
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Suppose we have a function call f(s).
	// Here’s how we’d call that in the usual way, running it synchronously.
	f("direct")

	// To invoke this function in a goroutine, use go f(s).
	// This new goroutine will execute concurrently with the calling one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in separate goroutines now.
	// Wait for them to finish (for a more robust approach, use a WaitGroup).
	time.Sleep(time.Second)
	fmt.Println("done")
}
```

When we run this program, we see the output of the blocking call first, then the output of the two goroutines. The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.

```sh
go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done
```

## Channels

_Channels_ are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.

```go
package main

import "fmt"

func main() {
	// Create a new channel with make(chan val-type).
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	// Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() {
		messages <- "ping"
	}()

	// The <-channel syntax receives a value from the channel.
	// Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
}
```

When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.

```sh
go run channels.go
ping
```

## Channel buffering

By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.

```go
package main

import "fmt"

func main() {
	// Here we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
```

```sh
go run channel-buffering.go
buffered
channel
```

## Channel Synchronization

We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish.
When waiting for multiple goroutines to finish, you may prefer to use a _WaitGroup_.

```go
package main

import (
	"fmt"
	"time"
)

// This is the function we’ll run in a goroutine.
// The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we’re done.
	done <- true
}

func main() {
	// Start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
}
```

```sh
go run channel-synchronization.go
working...done
```

If you removed the <- done line from this program, the program could exit before the worker finished its work, or in some cases even before it started.

## Channel Directions

When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
This specificity increases the type-safety of the program.

```go
package main

import "fmt"

// This ping function only accepts a channel for sending values.
// It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
```

```sh
go run channel-directions.go
passed message
```

## Select

Go’s _select_ lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful feature of Go.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// For our example we’ll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time,
	// to simulate e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	fmt.Println("Elapsed:", time.Since(start))
}
```

```sh
time go run select.go
received one
received two
Elapsed: 2.00111067s
```

## Timeouts

_Timeouts_ are important for programs that connect to external resources or that otherwise need to bound execution time.
Implementing timeouts in Go is easy and elegant thanks to channels and select.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// For our example, suppose we’re executing an external call that returns its result on a channel c1 after 2s.
	// Note that the channel is buffered, so the send in the goroutine is nonblocking.
	// This is a common pattern to prevent goroutine leaks in case the channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Here’s the select implementing a timeout.
	// res := <-c1 awaits the result and <-time.
	// After awaits a value to be sent after the timeout of 1s.
	// Since select proceeds with the first receive that’s ready,
	// we’ll take the timeout case if the operation takes more than the allowed 1s.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we’ll print the result.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
```

Running this program shows the first operation timing out and the second succeeding.

```sh
go run timeouts.go
timeout 1
result 2
```

## Non-Blocking Channel Operations

Basic sends and receives on channels are blocking.
However, we can use select with a default clause to implement _non-blocking_ sends, receives, and even non-blocking multi-way selects.

```go
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here’s a non-blocking receive.
	// If a value is available on messages then select will take the <-messages case with that value.
	// If not it will immediately take the default case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly.
	// Here msg cannot be sent to the messages channel,
	// because the channel has no buffer and there is no receiver.
	// Therefore the default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple cases above the default clause to implement a multi-way non-blocking select.
	// Here we attempt non-blocking receives on both messages and signals.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
```

```sh
go run 035_non_block/non_blocking.go
no message received
no message sent
no activity
```

## Closing Channels

_Closing_ a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.

```go
package main

import "fmt"

// In this example we’ll use a jobs channel to communicate work to be done
// from the main() goroutine to a worker goroutine.
// When we have no more jobs for the worker we’ll close the jobs channel.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here’s the worker goroutine.
	// It repeatedly receives from jobs with j, more := <-jobs.
	// In this special 2-value form of receive,
	// the more value will be false if jobs has been closed
	// and all values in the channel have already been received.
	// We use this to notify on done when we’ve worked all our jobs.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// This sends 3 jobs to the worker over the jobs channel, then closes it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the synchronization approach we saw earlier.
	<-done

	// Reading from a closed channel succeeds immediately,
	// returning the zero value of the underlying type.
	// The optional second return value is true
	// if the value received was delivered by a successful send operation to the channel,
	// or false if it was a zero value generated because the channel is closed and empty.
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
```

```sh
go run 036_close_ch/closing_channel.go
sent job 1
received job 1
received job 2
sent job 2
sent job 3
sent all jobs
received job 3
received all jobs
received more jobs: false
```

## Range over Channels

In a previous example we saw how for and range provide iteration over basic data structures.
We can also use this syntax to iterate over values received from a channel

```go
package main

import "fmt"

func main() {
	// We’ll iterate over 2 values in the queue channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This range iterates over each element as it’s received from queue.
	// Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}

}
```

```sh
go run range-over-channels.go
one
two
```

## Timers

We often want to execute Go code at some point in the future, or repeatedly at some interval.
Go’s built-in timer and ticker features make both of these tasks easy.
We’ll look first at timers and then at tickers.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Timers represent a single event in the future.
	// You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
	// This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// If you just wanted to wait, you could have used time.Sleep.
	// One reason a timer may be useful is that you can cancel the timer before it fires.
	// Here’s an example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.
	time.Sleep(2 * time.Second)
}
```

```sh
go run timers.go
Timer 1 fired
Timer 2 stopped
```

The first timer will fire ~2s after we start the program, but the second should be stopped before it has a chance to fire.

## Tickers

Timers are for when you want to do something once in the future
Tickers are for when you want to do something repeatedly at regular intervals.
Here’s an example of a ticker that ticks periodically until we stop it.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Tickers use a similar mechanism to timers: a channel that is sent values.
	// Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers can be stopped like timers.
	// Once a ticker is stopped it won’t receive any more values on its channel.
	// We’ll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
```

```sh
go run tickers.go
Tick at 2012-09-23 11:29:56.487625 -0700 PDT
Tick at 2012-09-23 11:29:56.988063 -0700 PDT
Tick at 2012-09-23 11:29:57.488076 -0700 PDT
Ticker stopped
```

When we run this program the ticker should tick 3 times before we stop it.

## Worker pools

In this example we’ll look at how to implement a _worker pool_ using goroutines and channels.

```go
package main

import (
	"fmt"
	"time"
)

// Here’s the worker, of which we’ll run several concurrent instances.
// These workers will receive work on the jobs channel and send the corresponding results on results.
// We’ll sleep a second per job to simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// In order to use our pool of workers we need to send them work and collect their results.
	// We make 2 channels for this.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 jobs and then close that channel to indicate that’s all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have finished.
	// An alternative way to wait for multiple goroutines is to use a WaitGroup.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
```

```sh
time go run worker-pools.go
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5
real    0m2.358s
```

Our running program shows the 5 jobs being executed by various workers.
The program only takes about 2 seconds despite doing about 5 seconds of total work because there are 3 workers operating concurrently.

## WaitGroups

To wait for multiple goroutines to finish, we can use a _wait group_.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we’ll run in every goroutine.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This WaitGroup is used to wait for all the goroutines launched here to finish.
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	var wg sync.WaitGroup

	// Launch several goroutines using WaitGroup.Go
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}

	// Block until all the goroutines started by wg are done.
	// A goroutine is done when the function it invokes returns.
	wg.Wait()
}
```

```sh
go run waitgroups.go
Worker 5 starting
Worker 3 starting
Worker 4 starting
Worker 1 starting
Worker 2 starting
Worker 4 done
Worker 1 done
Worker 2 done
Worker 5 done
Worker 3 done
```

Note that this approach has no straightforward way to propagate errors from workers. For more advanced use cases, consider using the [errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup).

## Rate Limiting

[Rate limiting](https://en.wikipedia.org/wiki/Rate_limiting) is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and tickers.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// First we’ll look at basic rate limiting.
	// Suppose we want to limit our handling of incoming requests.
	// We’ll serve these requests off a channel of the same name.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds.
	// This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the limiter channel before serving each request,
	// we limit ourselves to 1 request every 200 milliseconds.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// We may want to allow short bursts of requests in our rate limiting scheme
	// while preserving the overall rate limit.
	// We can accomplish this by buffering our limiter channel.
	// This burstyLimiter channel will allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requests.
	// The first 3 of these will benefit from the burst capability of burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
```

```sh
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC
```

## Atomic Counters


The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools. There are a few other options for managing state though. Here we’ll look at using the *sync/atomic* package for *atomic counters* accessed by multiple goroutines.

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// We’ll use an atomic integer type to represent our (always-positive) counter.
	var ops atomic.Uint64
	//A WaitGroup will help us wait for all goroutines to finish their work.
	var wg sync.WaitGroup

	// We’ll start 50 goroutines that each increment the counter exactly 1000 times.
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				// To atomically increment the counter we use Add.
				ops.Add(1)
			}
		})
	}

	// Wait until all the goroutines are done.
	wg.Wait()

	// Here no goroutines are writing to ‘ops’,
	// but using Load it’s safe to atomically read a value even while other goroutines are (atomically) updating it.
	fmt.Println("ops:", ops.Load())
}
```

```sh
go run atomic-counters.go
ops: 50000
```

We expect to get exactly 50,000 operations. Had we used a non-atomic integer and incremented it with ops++, we’d likely get a different number, changing between runs, because the goroutines would interfere with each other. Moreover, we’d get data race failures when running with the -race flag.

## Mutexes

In the previous example we saw how to manage simple counter state using atomic operations. For more complex state we can use a [mutex](https://en.wikipedia.org/wiki/Mutual_exclusion) to safely access data across multiple goroutines.

```go
package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters;
// since we want to update it concurrently from multiple goroutines,
// we add a Mutex to synchronize access.
// Note that mutexes must not be copied, so if this struct is passed around,
// it should be done by pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Lock the mutex before accessing counters;
	// unlock it at the end of the function using a defer statement.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// Note that the zero value of a mutex is usable as-is,
		// so no initialization is required here.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// This function increments a named counter in a loop.
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// Run several goroutines concurrently;
	// note that they all access the same Container,
	// and two of them access the same counter.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("a", 10000)
	})

	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// Wait for the goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)
}
```

```sh
go run mutexes.go
map[a:20000 b:10000]
```
