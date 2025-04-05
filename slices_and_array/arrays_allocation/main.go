package main

// go build -gcflags='-m' main.go | grep escape; rm main.exe || main

func allocation() *[3]int8 {
	var data [3]int8
	return &data // moved to heap
}

func main() {
	var arrayWithStack [128 << 10]int8 // 128 KB
	_ = arrayWithStack

	var arrayWithHeap [12 << 20]int8 // 12 MB // moved to heap
	_ = arrayWithHeap

	arrayWithHeap2 := allocation() // in heap
	_ = arrayWithHeap2
}
