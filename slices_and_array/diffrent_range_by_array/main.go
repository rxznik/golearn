package main

import "fmt"

func main() {
	data := [...]int{1, 2, 3}

	for _, value := range data { // copy of array
		fmt.Print(value)
	}
	fmt.Println()

	for _, value := range &data { // copy of pointer
		fmt.Print(value)
	}
	fmt.Println()

	for _, value := range data[:] { // copy of slice
		fmt.Print(value)
	}
	fmt.Println()
}
