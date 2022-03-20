package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	m1 := map[int]string{
		123: "foo",
		456: "bar",
	}
	fmt.Println(maps.Keys(m1))   // [123 456]
	fmt.Println(maps.Values(m1)) // [foo bar]
	fmt.Println("m1", m1)

	m2 := maps.Clone(m1)
	fmt.Println("m2", m2)
	fmt.Println(maps.Equal(m2, m1))

	m3 := map[int]string{
		789: "baz",
	}

	maps.Copy(m3, m1)
	fmt.Println("m3", m3)
	fmt.Println(maps.Equal(m3, m1))
}
