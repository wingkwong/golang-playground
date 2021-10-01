package main

import (
	"fmt"
)

func main() {
	// simple case
	var x bool = false
	fmt.Printf("%v, %T\n", x, x)

	// logical test
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	// default value is false
	var y bool
	fmt.Printf("%v, %T\n", y, y)
}
