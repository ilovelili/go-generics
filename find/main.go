package main

import "fmt"

func Find[T comparable](src []T, target T) int {
	for i, item := range src {
		if item == target {
			return i
		}
	}

	return -1
}

func main() {
	src := []int{1, 2, 3}
	target := 3
	fmt.Println(Find(src, target))
}
