package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var data []string
	fmt.Println("\x1b[33mvar data []string\x1b[0m")
	fmt.Printf(
		"\tempty=%t nil=%t size=%d data=%p\n",
		len(data) == 0,
		data == nil,
		unsafe.Sizeof(data),
		unsafe.SliceData(data),
	) // массив не проинициализирован, nil

	data = []string(nil)
	fmt.Println("\x1b[33mdata = []string(nil)\x1b[0m")
	fmt.Printf(
		"\tempty=%t nil=%t size=%d data=%p\n",
		len(data) == 0,
		data == nil,
		unsafe.Sizeof(data),
		unsafe.SliceData(data),
	) // массив не проинициализирован, nil

	data = make([]string, 0)
	fmt.Println("\x1b[33mdata = make([]string, 0)\x1b[0m")
	fmt.Printf(
		"\tempty=%t nil=%t size=%d data=%p\n",
		len(data) == 0,
		data == nil,
		unsafe.Sizeof(data),
		unsafe.SliceData(data),
	) // массив проинициализирован, не nil

	empty := struct{}{}
	fmt.Println("\x1b[33mempty := struct{}{}\x1b[0m")
	fmt.Printf("\tsize=%d empty=%p\n", unsafe.Sizeof(empty), unsafe.Pointer(&empty))
}
