# [Go by examples](https://gobyexample.com/)

[Go](https://go.dev/) is an open source programming language designed for building scalable, secure and reliable software. Please read the [official documentation](https://go.dev/doc/tutorial/getting-started) to learn more.

_Go by Example_ is a hands-on introduction to Go using annotated example programs. Check out the [first example](https://gobyexample.com/hello-world) or browse the full list below.

Unless stated otherwise, examples here assume the [latest major release Go](https://go.dev/doc/devel/release) and may use new language features. Try to upgrade to the latest version if something isn't working.

## 1.Hello world

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

## 2.Values

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

## 3.Variables

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

## 4.Constants

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

## 5.For

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

## 6.If

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

## 7.Switch

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

## 8.Arrays

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

## 9.Slices

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

## 10.Maps

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

## 11.Functions

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

## 12.Multiple Return Values

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

## 13.Variadic functions

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

## 14.Closures

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

## 15.Recursion

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

## 16.Range over Built-in Types

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

## 17.Pointers

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

## 18.Strings and runes

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

## 19.Structs

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

## 20.Methods

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

## 21.Interfaces

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

## 22.Enums

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

## 23.Struct Embedding

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

## Generics

Starting with version 1.18, Go has added support for generics.
More info in [this](https://go.dev/blog/deconstructing-type-parameters) post and in stdlib [docs](https://pkg.go.dev/slices#Index)

```go
package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}

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

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
    var s = []string{"foo", "bar", "zoo"}

    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

    _ = SlicesIndex[[]string, string](s, "zoo")

    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())
}
```

## Range over Iterators

Starting with version 1.23, Go has added support for [iterators](https://go.dev/blog/range-functions), which lets us range over pretty much anything.
The iterator function takes another function as a parameter, called yield by convention.
Packages like [slices](https://pkg.go.dev/slices) have a number of useful functions to work with iterators.

```go
package main

import (
    "fmt"
    "iter"
    "slices"
)

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

func (lst *List[T]) All() iter.Seq[T] {
    return func(yield func(T) bool) {

        for e := lst.head; e != nil; e = e.next {
            if !yield(e.val) {
                return
            }
        }
    }
}

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

    for e := range lst.All() {
        fmt.Println(e)
    }

    all := slices.Collect(lst.All())
    fmt.Println("all:", all)

    for n := range genFib() {

        if n >= 10 {
            break
        }
        fmt.Println(n)
    }
}
```

## Errors

In Go it’s idiomatic to communicate [errors](https://pkg.go.dev/errors) via an explicit, separate return [value](https://go.dev/blog/go1.13-errors). This contrasts with the exceptions.
By convention, errors are the last return value and have type error, a built-in interface.
A nil value in the error position indicates that there was no error.
A sentinel error is a predeclared variable that is used to signify a specific error condition.
_errors.Is_ checks that a given error.

```go
package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

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
