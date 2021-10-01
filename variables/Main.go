package main

import "fmt"

var foo = "foo"
var bar = "bar"
var baz = "baz"

var (
	foo2 = "foo"
	bar2 = "bar"
	baz2 = "baz"
)

func main() {
	// You should merge variable declaration with assignment on next line
	// var i int
	// i = 999
	var i int = 999

	// use := for implicit type
	j := 888
	k := true

	// It should print : 999 888 true
	fmt.Println(i, j, k)

	// Print the types
	fmt.Printf("%T, %T, %T\n", i, j, k)

	// Print the values
	fmt.Printf("%v, %v, %v\n", i, j, k)
	fmt.Println(foo, bar, baz)
	fmt.Println(foo2, bar2, baz2)

	n := 0
	if true {
		n := 1 // Variable Shadowing !
		// You should write
		// n = 1
		// instead
		n++
	}
	fmt.Println(n) // 0
}
