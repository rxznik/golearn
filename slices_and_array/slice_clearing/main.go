package main

import "fmt"

func main() {
	// remove via nil
	first := []int{1, 2, 3, 4, 5}
	first = nil
	fmt.Printf("first: %v, len: %d, cap: %d\n", first, len(first), cap(first))

	// remove via slice
	second := []int{1, 2, 3, 4, 5}
	second = second[:0]
	fmt.Printf("second: %v, len: %d, cap: %d\n", second, len(second), cap(second))

	// remove via make
	third := []int{1, 2, 3, 4, 5}
	third = make([]int, 0, cap(third))
	fmt.Printf("third: %v, len: %d, cap: %d\n", third, len(third), cap(third))

	// remove via clear
	fourth := []int{1, 2, 3, 4, 5}
	clear(fourth)
	fmt.Printf("fourth: %v, len: %d, cap: %d\n", fourth, len(fourth), cap(fourth))

	// zeroing n-elements via clear
	fifth := []int{1, 2, 3, 4, 5}
	clear(fifth[1:3])
	fmt.Printf("fifth: %v, len: %d, cap: %d\n", fifth, len(fifth), cap(fifth))
}
