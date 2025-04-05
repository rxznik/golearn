package main

import (
	"fmt"
	"runtime"
)

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%.3f MB\n", float32(m.Alloc)/(1<<20))
}

func main() {
	slice := make([]int8, 0, 3)
	fmt.Printf("slice: %v - slice address: %p\n", slice, &slice)
	printAllocs() // 0 MB

	slice = append(slice, 1, 2, 3)
	fmt.Println("slice full capacity:", slice, "len:", len(slice), "cap:", cap(slice))
	printAllocs() // 0 MB

	slice = append(slice, make([]int8, 60<<10)...)
	fmt.Printf("slice address: %p\n", &slice)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	printAllocs() // 0 MB

	//// проверка ф-ции printAllocs
	slice = make([]int8, 1<<20)
	printAllocs()

	runtime.GC()
	printAllocs()
}
