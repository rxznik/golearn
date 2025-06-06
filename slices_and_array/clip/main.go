package main

import (
	"fmt"
	"runtime"
	"slices"
	"unsafe"
)

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/(1<<20))
}

func main() {
	data := make([]int, 10, 100<<20)
	fmt.Println("data", unsafe.SliceData(data), len(data), cap(data))

	printAllocs()

	temp1 := data[:10]
	fmt.Println("temp1", unsafe.SliceData(temp1), len(temp1), cap(temp1))

	temp2 := slices.Clip(data) // data[:len(data):len(data)]
	fmt.Println("temp2", unsafe.SliceData(temp2), len(temp2), cap(temp2))

	runtime.GC()
	printAllocs()

	runtime.KeepAlive(temp1)
	runtime.KeepAlive(temp2)
	/// если нужна лишь малая часть среза, то нужно копировать,
	/// чтобы gc мог почистить память от большего среза
}
