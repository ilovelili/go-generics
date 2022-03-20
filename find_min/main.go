package main

import (
	"golang.org/x/exp/constraints"
)

type ElementType interface {
	int | string
}

func Min1[T ElementType](t1, t2 T) T {
	if t1 > t2 {
		return t2
	}

	return t1
}

func Min2[T constraints.Ordered](t1, t2 T) T {
	if t1 > t2 {
		return t2
	}

	return t1
}

func main() {
	println(Min1(1, 2))
	println(Min2("2", "1"))
}
