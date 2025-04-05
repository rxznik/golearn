package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4}
	fmt.Printf("data: %v, len: %d, cap: %d\n", data, len(data), cap(data))

	process1(data)
	fmt.Printf("data: %v, len: %d, cap: %d\n", data, len(data), cap(data))

	process2(data)
	fmt.Printf("data: %v, len: %d, cap: %d\n", data, len(data), cap(data))
}

func process1(data []int) {
	if len(data) == 0 {
		return
	}
	data[0] = 1488
}

func process2(data []int) {
	data = append(data, 1488)
	fmt.Printf("\x1b[35mdata: %v, len: %d, cap: %d\x1b[0m\n", data, len(data), cap(data))
}
