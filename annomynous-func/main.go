package main

func compare[T comparable]() func(T, T) bool {
	return func(t1, t2 T) bool {
		return t1 == t2
	}
}

func main() {
	c := compare[int64]()
	println(c(1, 1))
}
