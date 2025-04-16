package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	tickets chan struct{}
}

func NewSemaphore(numTickets int) Semaphore {
	return Semaphore{tickets: make(chan struct{}, numTickets)}
}

func (s *Semaphore) Acquire() {
	s.tickets <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.tickets
}

const numOperations = 12

func main() {
	var wg sync.WaitGroup
	wg.Add(numOperations)

	semaphore := NewSemaphore(3)
	for i := range numOperations {
		semaphore.Acquire()
		go func(idx int) {
			defer func() {
				wg.Done()
				semaphore.Release()
			}()
			fmt.Println(idx, "working...")
			time.Sleep(1 * time.Second)
			fmt.Println(idx, "exiting...")
		}(i)
	}

	wg.Wait()
}
