package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	a := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
		"END",
	}

	a = a[:12]
	fmt.Println("a", a)
	fmt.Println(cap(a)) // 13

	a = slices.Clip(a)
	fmt.Println("a", a)
	fmt.Println(cap(a)) // 12

	b := []string{
		"foo",
		"foo",
		"bar",
		"baz",
		"foo",
		"baz",
	}

	// compact only replaces consecutive runs of equal elements
	a = slices.Compact(b) // [foo bar baz foo baz]

	fmt.Println(slices.Contains(a, "zoo")) // false

	a = slices.Delete(a, 1, 3)
	fmt.Println(a) // [foo foo baz]

	fmt.Println("a", a)
	fmt.Println(cap(a))

	a = slices.Grow(a, 5)
	fmt.Println("a growed to", a)
	fmt.Println(cap(a))

	a = slices.Insert(a, 0, "noo")
	fmt.Println("a", a)
	fmt.Println(cap(a))
}
