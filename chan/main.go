package main

import (
	"context"
)

type ChanElement interface {
	int | string
}

func makeChan[T chan E, E ChanElement](ctx context.Context, arr []E) T {
	ch := make(chan E)

	go func() {
		// close the channel when all received
		defer close(ch)
		for _, v := range arr {
			select {
			case <-ctx.Done():
				return
			default:
			}

			// time.Sleep(time.Duration(rand.Intn(3) * int(time.Second)))
			ch <- v
		}
	}()

	return ch
}

func main() {
	for v := range makeChan(context.Background(), []int{1, 2, 3}) {
		println(v)
	}

	for v := range makeChan(context.Background(), []string{"1", "2", "3"}) {
		println(v)
	}
}
