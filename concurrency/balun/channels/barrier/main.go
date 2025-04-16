package main

import (
	"fmt"
	"sync"
)

type Barrier struct {
	count int
	size  int
	mutex sync.Mutex

	beforeCh chan struct{}
	afterCh  chan struct{}
}

func NewBarrier(size int) *Barrier {
	return &Barrier{
		size:     size,
		beforeCh: make(chan struct{}, size),
		afterCh:  make(chan struct{}, size),
	}
}

func (b *Barrier) Before() {
	b.mutex.Lock()

	b.count++

	if b.count == b.size {
		for range b.size {
			b.beforeCh <- struct{}{}
		}
	}

	b.mutex.Unlock()
	<-b.beforeCh
}

func (b *Barrier) After() {
	b.mutex.Lock()
	b.count--

	if b.count == 0 {
		for range b.size {
			b.afterCh <- struct{}{}
		}
	}

	b.mutex.Unlock()
	<-b.afterCh
}

const numOperations = 3

func main() {
	var wg sync.WaitGroup
	wg.Add(numOperations)

	barrier := NewBarrier(numOperations)
	for i := range numOperations {
		go func(idx int) {
			defer wg.Done()
			for i := range 10 {
				fmt.Printf("goroutine %d: working with %d\n", idx+1, i+1)

				barrier.Before() // blocked, while waiting for all goroutines
				fmt.Printf("\x1b[33mgoroutine %d: bootstrap with %d\x1b[0m\n", idx+1, i+1)
				barrier.After() // blocked, while all goroutines are working with this part ^

				fmt.Printf("goroutine %d: finished with %d\n", idx+1, i+1)
			}
		}(i)
	}

	wg.Wait()
}
