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
