package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func sort[E constraints.Ordered](arr []E) []E {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(sort([]int{1, 3, 2}))
	fmt.Println(sort([]string{"1", "3", "2"}))
}
