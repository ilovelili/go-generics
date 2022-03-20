package main

import (
	"golang.org/x/exp/maps"
)

func main() {
	m1 := map[int]string{
		123: "foo",
		456: "bar",
	}
	println(maps.Keys(m1))   // [123 456]
	println(maps.Values(m1)) // [foo bar]
	println("m1", m1)

	m2 := maps.Clone(m1)
	println("m2", m2)
	println(maps.Equal(m2, m1))

	m3 := map[int]string{
		789: "baz",
	}

	maps.Copy(m3, m1)
	println("m3", m3)
	println(maps.Equal(m3, m1))
}
