package main

import (
	"fmt"
	"unsafe"
)

func Append[T any](slice []T, elems ...T) []T {
	previousLength := len(slice)
	newLength := previousLength + len(elems)
	if newLength > cap(slice) {
		capacity := newLength * 2
		if capacity == 0 {
			capacity = 1
		}
		newSlice := make([]T, newLength, capacity)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[:newLength]
	copy(slice[previousLength:newLength], elems)
	return slice
}

func main() {
	data := make([]int, 0, 2)
	fmt.Println(data, "cap:", cap(data), "addr: %p", unsafe.SliceData(data))

	data = Append(data, 1, 2)
	fmt.Println(data, "cap:", cap(data), "addr: %p", unsafe.SliceData(data))

	data = Append(data, 3)
	fmt.Println(data, "cap:", cap(data), "addr: %p", unsafe.SliceData(data))

	data = Append(data, 4, 5, 6)
	fmt.Println(data, "cap:", cap(data), "addr: %p", unsafe.SliceData(data))
}
