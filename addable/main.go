package main

import (
	"fmt"
	"math"
)

type Addable interface {
	// https://stackoverflow.com/questions/70888240/whats-the-meaning-of-the-new-token-tilde-operator-in-go-approximation-elem
	// tilde token allows underlying types
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// generic type
type Foo[T Addable] struct {
	v T
}

func (foo Foo[T]) Add(v T) {
	foo.v += v
}

type MyInt int

func main() {
	f1 := new(Foo[float64])
	f1.v = math.Pi

	f1.Add(1.1)
	fmt.Println(f1.v)

	f2 := new(Foo[MyInt]) // Addable must be ~int otherwise the underlying type MyInt will cause compiling error
	f2.Add(3)
	fmt.Println(f1.v)
}
