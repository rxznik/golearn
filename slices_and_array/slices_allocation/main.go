package main

// go build -gcflags='-m' main.go | grep escape; rm main.exe || main

func allocation() *[]int8 {
	data := make([]int8, 1<<10)
	return &data
}

func main() {
	sliceWithStack := make([]int8, 0, 64<<10) // 64 KB
	_ = sliceWithStack

	var arrayWithStack [128 << 10]int8
	sliceWithStack2 := arrayWithStack[:] // 128 KB
	_ = sliceWithStack2

	sliceWithHeap := make([]int8, (64<<10)+1) // 65 KB // moved to heap
	_ = sliceWithHeap

	sliceInHeap := allocation() // in heap
	_ = sliceInHeap
}
