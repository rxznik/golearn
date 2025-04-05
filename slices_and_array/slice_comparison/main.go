package main

import (
	"fmt"
	"reflect"
)

func main() {
	data1 := []int{1, 2, 3, 4, 5}
	data2 := []int{1, 2, 3, 4, 5}

	//// compilation error
	// if data1 == data2 {
	// 	fmt.Println("equal")
	// }

	isEqual := reflect.DeepEqual(data1, data2)
	fmt.Println("equal:", isEqual)
}
