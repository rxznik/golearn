package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice := make([]int8, 0, 3)
	fmt.Printf("slice: %v\nslice address: %p\n", slice, unsafe.SliceData(slice))

	slice = append(slice, 1, 2, 3)
	fmt.Println("slice full capacity:", slice, "len:", len(slice), "cap:", cap(slice))
	fmt.Println("slice address:", unsafe.SliceData(slice)) // адресс остался неизменен

	/// превосхождение capacity
	slice = append(slice, 4)
	/// адресс массива сильно изменился, произошла аллокация в кучу
	fmt.Printf("slice: %v\nslice address: %p\n", slice, unsafe.SliceData(slice))
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}
