package main

import "fmt"

func main() {
	slice := make([]int, 3, 6)
	array := [3]int(slice[:3])

	slice[0] = 10
	fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))
	fmt.Printf("array: %v, len: %d, cap: %d\n", array, len(array), cap(array))

	slice2 := make([]int, 3, 6)
	array2 := (*[3]int)(slice2[:3])

	slice2[0] = 10
	fmt.Printf("slice2: %v, len: %d, cap: %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("array2: %v, len: %d, cap: %d\n", array2, len(array2), cap(array2))
}
