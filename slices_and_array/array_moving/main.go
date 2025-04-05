package main

import (
	"fmt"
	"unsafe"
)

func allocation(index int) byte {
	//// стек в Go равен 8 KB, придётся увеличивать стек
	var data [1 << 20]byte // 1 MB, т.к. 2^20 / 1024^2 = 1 MB
	return data[index]
}

func main() {
	var array [10]int
	address := (uintptr)(unsafe.Pointer(&array))
	fmt.Printf("First address: %x\n", address)

	allocation(100)

	address = (uintptr)(unsafe.Pointer(&array))
	fmt.Printf("Second address: %x\n", address)
}
