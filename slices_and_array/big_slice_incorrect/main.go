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

func FindData(filename string) []byte {
	data := make([]byte, 100<<20)

	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 {
			return data[i : i+20]
		}
	}

	return nil
}

func main() {
	data := FindData("filename.bin")
	printAllocs()

	runtime.GC()
	runtime.KeepAlive(data)
	printAllocs()
}
