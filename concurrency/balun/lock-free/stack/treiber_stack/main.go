package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

type Item struct {
	value int
	next  unsafe.Pointer
}

type Stack struct {
	head unsafe.Pointer
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(value int) {
	item := &Item{value, atomic.LoadPointer(&s.head)}
	for {
		head := atomic.LoadPointer(&s.head)
		item.next = head
		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(item)) {
			return
		}
	}
}

func (s *Stack) Pop() int {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return 0
		}

		next := atomic.LoadPointer(&(*Item)(head).next)
		if atomic.CompareAndSwapPointer(&s.head, head, next) {
			return (*Item)(head).value
		}
	}
}

func (s *Stack) Peek() int {
	head := atomic.LoadPointer(&s.head)
	if head == nil {
		return 0
	}

	return (*Item)(head).value
}

const operationsCount = 3

func main() {
	stack := NewStack()

	var wg sync.WaitGroup
	wg.Add(operationsCount)

	for i := range operationsCount {
		go func(value int) {
			defer wg.Done()
			stack.Push(value * 100)
			fmt.Println(value, "after push:", stack.Peek())
		}(i)

		go func(value int) {
			defer wg.Done()
			fmt.Println(value, "popped:", stack.Pop())
			fmt.Println(value, "after pop:", stack.Peek())
		}(i)
	}

	wg.Wait()
}
