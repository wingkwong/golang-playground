# Variables

Define a variable

```go
var i int
```

Assign a value

```go
i = 42
```

However, you should merge variable declaration with assignment on next line

```go
func main() {
    var i int = 42
}
```

Inside a function, you can replace ``var`` with ``:=`` wiuth implicit type.

```go
func main() {
    i := 42
}
```

Use ``fmt.Println`` to print the values

```go
func main() {
    fmt.Println(i, j, k)
}
```

You should see

```
999 888 true
```

You can also check the type using ``fmt.Printf``

```go
func main() {
    fmt.Printf("%T, %T, %T\n", i, j, k)
}
```

of course you can also use it to check the values

```go
func main() {
    fmt.Printf("%v, %v, %v\n", i, j, k)
}
```

We can also define variables outside a function (i.e. in a package scope)

```go
var foo = "foo"
var bar = "bar"
var baz = "baz"

func main() {
	fmt.Println(foo, bar, baz)
}
```

We can wrap the whole section with a ``var`` block

```go
var (
    foo = "foo"
    bar = "bar"
    baz = "baz"
)

func main() {
	fmt.Println(foo, bar, baz)
}
```

Let's take a look at the following code. You should see ``n`` would be 0. 

```go
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n) // 0
```

This is because the statement ``n := 1`` declares a new variables ``n`` shadowing the original one in the scope of the if statement. If we need to reuse ``n`` from outer block, we should write the following instead.

```go
func main() {
	n := 0
	if true {
		n = 1
		n++
	}
	fmt.Println(n) // 2
```