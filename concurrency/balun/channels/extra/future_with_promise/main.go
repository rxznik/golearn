package main

import (
	"fmt"
	"time"
)

type Future[T any] struct {
	resultCh <-chan T
}

func NewFuture[T any](resultCh <-chan T) Future[T] {
	return Future[T]{
		resultCh: resultCh,
	}
}

func (f *Future[T]) Get() T {
	return <-f.resultCh
}

type Promise[T any] struct {
	resultCh chan T
}

func NewPromise[T any]() Promise[T] {
	return Promise[T]{resultCh: make(chan T)}
}

func (p *Promise[T]) Set(value T) {
	p.resultCh <- value
	close(p.resultCh)
}

func (p *Promise[T]) GetFuture() Future[T] {
	return NewFuture(p.resultCh)
}

func main() {
	promise := NewPromise[string]()

	go func() {
		time.Sleep(time.Second)
		promise.Set("done")
	}()

	future := promise.GetFuture()
	value := future.Get()
	fmt.Println("value:", value)
}
