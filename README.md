[![Go Reference][pkg.go.dev badge]][pkg.go.dev]
[![Go Report Card][goreportcard badge]][goreportcard]

# Curry

Curry provides facilities for [currying][currying] Go functions. The process
converts a single function like `f(x, y, z)` into N many functions each taking
a single argument like `h(x)(y)(z)`.

Currying can be useful as a way to provide partial function application. For
example, this takes a function that adds two integers and creates a new
function `plus1` that increments the given value by 1.

```go
package main

import (
  "fmt"

  "github.com/calebcase/curry"
)

func Add(x, y int) int {
  return x + y
}

func main() {
  // Create a reusable function with the first argument bound to 1.
  plus1 := curry.A2R1(Add)(1)

  fmt.Println(plus1(2)) // 3
  fmt.Println(plus1(3)) // 4
  fmt.Println(plus1(4)) // 5
}
```

For more examples and usage see the [go docs][pkg.go.dev].

## Build

The library is entirely generated with [jennifer][jennifer]:

```
go generate
go test -v
```

---

[currying]: https://en.wikipedia.org/wiki/Currying
[goreportcard badge]: https://goreportcard.com/badge/github.com/calebcase/curry
[goreportcard]: https://goreportcard.com/report/github.com/calebcase/curry
[jennifer]: https://github.com/dave/jennifer
[pkg.go.dev badge]: https://pkg.go.dev/badge/github.com/calebcase/curry.svg
[pkg.go.dev]: https://pkg.go.dev/github.com/calebcase/curry
