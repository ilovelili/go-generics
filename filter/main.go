package main

func filter[E any](arr []E, f func(E) bool) []E {
	result := []E{}
	for _, item := range arr {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 3}
	b := filter(a, func(i int) bool {
		return i < 2
	})

	println(b)
}
