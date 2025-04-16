package main

import (
	"fmt"
	"log"
	"time"
)

type result[T any] struct {
	val T
	err error
}

type Promise[T any] struct {
	resultCh chan result[T]
}

func NewPromise[T any](asyncFn func() (T, error)) Promise[T] {
	promise := Promise[T]{resultCh: make(chan result[T])}

	go func() {
		defer close(promise.resultCh)

		val, err := asyncFn()
		promise.resultCh <- result[T]{val, err}
	}()

	return promise
}

func (p *Promise[T]) Then(succesFn func(T), errorFn func(error)) {
	go func() {
		result := <-p.resultCh
		if result.err != nil {
			errorFn(result.err)
		} else {
			succesFn(result.val)
		}
	}()
}

func (p *Promise[T]) Await() (T, error) {
	result := <-p.resultCh
	return result.val, result.err
}

func main() {
	asyncJob := func() (string, error) {
		time.Sleep(1 * time.Second)
		return "ok", nil
	}

	promise := NewPromise[string](asyncJob)
	promise.Then(
		func(value string) {
			fmt.Println(value)
		},
		func(err error) {
			log.Fatal(err)
		},
	)

	for i := range 5 {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("i = ", i)
	}

	secondAsyncJob := func() (string, error) {
		time.Sleep(1 * time.Second)
		return "done", nil
	}
	promise = NewPromise[string](secondAsyncJob)
	val, err := promise.Await()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
}
