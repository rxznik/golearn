package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	//// default way
	msg := "hello world, hello!"
	msg = strings.Replace(msg, "hello", "hi", 2)
	fmt.Println(msg)

	//// other way
	msg = "hello world, hello!"
	data := []byte(msg)
	pointer := unsafe.SliceData(data)
	msg = unsafe.String(pointer, len(msg))
	fmt.Println("initial:", msg)
	data[0] = 'H'
	fmt.Println("changed:", msg)
}
