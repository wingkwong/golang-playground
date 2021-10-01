# Primitives

## Boolean type

```go
// simple case
var x bool = false
fmt.Printf("%v, %T\n", x, x)
```

```go
// logical test
n := 1 == 1
m := 1 == 2
fmt.Printf("%v, %T\n", n, n)
fmt.Printf("%v, %T\n", m, m)
```

```go
// default value is false
var y bool
fmt.Printf("%v, %T\n", y, y)
```

## Numberic types

| type 	| range (from) 	| range (to) 	|
|---	|---	|---	|
| int8 	| -128 	| 127 	|
| int16 	| -32768 	| 32767 	|
| int32 	| -2147483648 	| 2147483647 	|
| int64 	| -9223372036854775808 	| 9223372036854775807 	| 

### Integers

```go
// simple case
var x uint16 = 42
fmt.Printf("%v, %T\n", x, x)
```

```go
// operators
a := 10 // (1010)
b := 3  // (0011)
fmt.Println(a + b)  // 13
fmt.Println(a - b)  // 7
fmt.Println(a * b)  // 30
fmt.Println(a / b)  // 3
fmt.Println(a % b)  // 1
fmt.Println(a & b)  // 2 (0010)
fmt.Println(a | b)  // 11 (1011)
fmt.Println(a ^ b)  // 9 (1001)
fmt.Println(a &^ b) // 8 (0100)
c := 8 // 2 ^ 3 
fmt.Println(c << 3) // 64 = 2 ^ 6 = 2 ^ 3 * 2 ^ 3
fmt.Println(c >> 3) // 1 = 2 ^ 3 / 2 ^ 3 = 2 ^ 0
```

```go
// invalid operation: d + e (mismatched types int and int8)
var d int = 10
var e int8 = 3
fmt.Println(d + e) 
```

### Floating point

```go
f := 3.14
f = 13.7e72
f = 2.1E14
fmt.Printf("%v, %T\n", f, f)

f1 := 10.2
f2 := 3.7
fmt.Println(f1 + f2) // 13.899999999999999
fmt.Println(f1 - f2) // 6.499999999999999
fmt.Println(f1 * f2) // 37.74
fmt.Println(f1 / f2) // 2.7567567567567566
```

### Complex numbers

```go
var complex_n complex64 = 1 + 2i
fmt.Printf("%v, %T\n", complex_n, complex_n) // (1+2i), complex64
fmt.Printf("%v, %T\n", real(complex_n), real(complex_n)) // 1, float32
fmt.Printf("%v, %T\n", imag(complex_n), imag(complex_n)) // 2, float32

var complex_n2 complex128 = complex(5, 12) 
fmt.Printf("%v, %T\n", complex_n2, complex_n2) // (5+12i), complex128
```

## Text types

```go
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
```