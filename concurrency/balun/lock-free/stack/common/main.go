package main

import (
	"fmt"
	"reflect"
)

type item[T any] struct {
	value T
	next  *item[T]
}

type Stack[T any] struct {
	head *item[T]
	size uint
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.head = &item[T]{value, s.head}
	s.size++
}

func (s *Stack[T]) Pop() T {
	if s.head == nil {
		typeOfValue := reflect.TypeOf((*T)(nil)).Elem()
		zeroValue := reflect.Zero(typeOfValue)
		return zeroValue.Interface().(T)
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value
}

func (s *Stack[T]) Peek() T {
	if s.head == nil {
		typeOfValue := reflect.TypeOf((*T)(nil)).Elem()
		zeroValue := reflect.Zero(typeOfValue)
		return zeroValue.Interface().(T)
	}
	return s.head.value
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Len() int {
	size := reflect.ValueOf(s.size).Int()
	return int(size)
}

func main() {
	stack := NewStack[int]()

	fmt.Println(stack.Len(), stack.Empty(), stack.Peek())

	stack.Push(1)
	stack.Push(2)

	fmt.Println(stack.Len(), stack.Empty(), stack.Peek())

	stack.Pop()

	fmt.Println(stack.Len(), stack.Empty(), stack.Peek())

	stack.Pop()

	fmt.Println(stack.Len(), stack.Empty(), stack.Peek())
}
