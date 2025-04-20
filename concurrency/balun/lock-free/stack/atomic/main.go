package main

import (
	"fmt"
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
	atomic.StorePointer(&s.head, unsafe.Pointer(item))
}

func (s *Stack) Pop() int {
	head := atomic.LoadPointer(&s.head)
	if head == nil {
		return 0
	}
	value := (*Item)(head).value
	next := atomic.LoadPointer(&(*Item)(head).next)
	atomic.StorePointer(&s.head, next)
	return value
}

func (s *Stack) Peek() int {
	head := atomic.LoadPointer(&s.head)
	if head == nil {
		return 0
	}

	return (*Item)(head).value
}

func main() {
	stack := NewStack()

	stack.Push(1)
	fmt.Println("head:", stack.Peek())

	stack.Push(2)
	fmt.Println("head:", stack.Peek())

	fmt.Println("popped:", stack.Pop())
	fmt.Println("head:", stack.Peek())

	fmt.Println("popped:", stack.Pop())
	fmt.Println("head:", stack.Peek())
}
