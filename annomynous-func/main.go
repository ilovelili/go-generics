package main

import "fmt"

func compare[T comparable]() func(T, T) bool {
	return func(t1, t2 T) bool {
		return t1 == t2
	}
}

func main() {
	c := compare[int64]()
	fmt.Println(c(1, 1))
}
