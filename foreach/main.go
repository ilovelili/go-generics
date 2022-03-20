package main

import "fmt"

func main() {
	each([]int{1, 2, 3}, func(i int) {
		fmt.Println(i)
	})
}

func each[E any](arr []E, f func(E)) {
	for _, item := range arr {
		f(item)
	}
}
