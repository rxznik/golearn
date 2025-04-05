package main

import (
	"fmt"
	"unsafe"
)

// // хоть slice является структурой, его zero value равен nil
type MySlice struct {
	Array unsafe.Pointer
	Len   int
	Cap   int
}

func main() {
	//// создание среза
	array := [3]int{1, 2, 3}
	slice1 := array[0:2] // [1 2]
	slice1 = append(slice1, 3)
	fmt.Println(slice1) // [1 2 3]

	slice2 := []int{1, 2, 3} // [1 2 3]
	fmt.Println(slice2)

	slice3 := []int{} // []
	fmt.Println(slice3)

	slice4 := make([]int, 3) // [0 0 0]
	fmt.Println(slice4)

	slice5 := []int{3: 4} // [0 0 4]
	fmt.Println(slice5)

	slice6 := []int{3: 4, 5, 1: 2} // [0 2 0 4 5]
	fmt.Println(slice6)

	slice7 := make([]int, 3, 5)
	//// [0 0 0] cap = 5, len = 3
	fmt.Printf("%v cap = %d, len = %d\n", slice7, cap(slice7), len(slice7))

	//// append
	slice := make([]int, 0, 5)
	slice = append(slice, 1)
	slice = append(slice, 2, 3)
	slice = append(slice, []int{4, 5}...)
	fmt.Println(slice) // [1 2 3 4 5]

	//// нулевой срез
	var slice8 []int // []
	//// так нельзя
	// slice8[0] = 1 // panic
	//// так можно
	slice8 = append(slice8, 1)
	fmt.Println(slice8) // [1]
	clear(slice8)       // []
	//// длина и ёмкость не изменились
	fmt.Printf("after clear: len = %d, cap = %d\n", len(slice8), cap(slice8))

	var data []int // []
	for _ = range data {
	} // ok
}
