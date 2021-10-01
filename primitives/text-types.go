package main

import (
	"fmt"
)

func main() {
	// simple case
	s := "hello world"
	fmt.Printf("%v, %T\n", s, s) // hello world, string

	// byte
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b) // [104 101 108 108 111 32 119 111 114 108 100], []uint8

	// rune
	r := 'a'
	fmt.Printf("%v, %T\n", r, r) // 97, int32

	var r2 rune = 'a'
	fmt.Printf("%v, %T\n", r2, r2) // 97, int32
}
