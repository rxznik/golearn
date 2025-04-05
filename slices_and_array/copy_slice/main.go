package main

import (
	"fmt"
	"slices"
)

func main() {
	/// classic way
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("copied:", dst)

	/// copy via std library
	dst2 := slices.Clone(src)
	fmt.Println("copied:", dst2)

	/// copy via append
	dst3 := append([]int(nil), src...)
	fmt.Println("copied:", dst3)
}
