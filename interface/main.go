package main

type Int64 interface {
	~int64 | ~uint64
}

type NumberPlayer[T Int64] interface {
	AddCounter() T
	MinusCounter() T
}

func NewNumberPlayer[T Int64]() NumberPlayer[T] {
	return &Counter[T]{
		t:       0,
		counter: 1,
	}
}

type Counter[T Int64] struct {
	t       T
	counter T
}

func (c Counter[T]) AddCounter() T {
	c.t += c.counter
	return c.t
}

func (c Counter[T]) MinusCounter() T {
	c.t -= c.counter
	return c.t
}

func main() {
	counter := NewNumberPlayer[int64]()
	counter.AddCounter()
	println(counter)
}
