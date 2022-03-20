package main

import (
	"golang.org/x/exp/constraints"
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
	println("a", a)
	println(cap(a)) // 13

	a = slices.Clip(a)
	println("a", a)
	println(cap(a)) // 12

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

	println(slices.Contains(a, "zoo")) // false

	a = slices.Delete(a, 1, 3)
	println(a) // [foo foo baz]

	println("a", a)
	println(cap(a))

	a = slices.Grow(a, 5)
	println("a growed to", a)
	println(cap(a))

	a = slices.Insert(a, 0, "noo")
	println("a", a)
	println(cap(a))

	slices.SortStableFunc(a, less[string])
	println(a)
}

func less[E constraints.Ordered](i, j E) bool {
	return i < j
}
