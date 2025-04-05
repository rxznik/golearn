package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const ELEM_SIZE = unsafe.Sizeof(int32(0)) // 4 байта
	array := [...]int32{1, 2, 3}
	//// получаем указатель на начало массива
	pointer := unsafe.Pointer(&array)
	fmt.Println(pointer) // <smth address>

	for idx := range array {
		//// получение элемента по индексу с помощью формулы:
		//// &elem = &array_start + idx * elem_size
		elem := *(*int32)(unsafe.Add(pointer, uintptr(idx)*ELEM_SIZE))
		fmt.Printf("array[%d] = %d\n", idx, elem)
	}

	dangerous1 := *(*int32)(unsafe.Add(pointer, uintptr(3)*ELEM_SIZE))
	dangerous2 := *(*int32)(unsafe.Add(pointer, uintptr(4)*ELEM_SIZE))
	fmt.Println("dangerous:", dangerous1, dangerous2)
}
