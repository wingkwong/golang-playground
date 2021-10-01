package main

import (
	"fmt"
)

func main() {
	// simple case
	var x uint16 = 42
	fmt.Printf("%v, %T\n", x, x)

	// operators
	a := 10             // (1010)
	b := 3              // (0011)
	fmt.Println(a + b)  // 13
	fmt.Println(a - b)  // 7
	fmt.Println(a * b)  // 30
	fmt.Println(a / b)  // 3
	fmt.Println(a % b)  // 1
	fmt.Println(a & b)  // 2 (0010)
	fmt.Println(a | b)  // 11 (1011)
	fmt.Println(a ^ b)  // 9 (1001)
	fmt.Println(a &^ b) // 8 (0100)
	c := 8              // 2 ^ 3
	fmt.Println(c << 3) // 64 = 2 ^ 6 = 2 ^ 3 * 2 ^ 3
	fmt.Println(c >> 3) // 1 = 2 ^ 3 / 2 ^ 3 = 2 ^ 0

	// invalid operation: d + e (mismatched types int and int8)
	// var d int = 10
	// var e int8 = 3
	// fmt.Println(d + e)

	// Floating point
	f := 3.14
	f = 13.7e72
	f = 2.1e14
	fmt.Printf("%v, %T\n", f, f)

	f1 := 10.2
	f2 := 3.7
	fmt.Println(f1 + f2) // 13.899999999999999
	fmt.Println(f1 - f2) // 6.499999999999999
	fmt.Println(f1 * f2) // 37.74
	fmt.Println(f1 / f2) // 2.7567567567567566

	// complex
	var complex_n complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", complex_n, complex_n)             // (1+2i), complex64
	fmt.Printf("%v, %T\n", real(complex_n), real(complex_n)) // 1, float32
	fmt.Printf("%v, %T\n", imag(complex_n), imag(complex_n)) // 2, float32

	var complex_n2 complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", complex_n2, complex_n2) // (5+12i), complex128
}
