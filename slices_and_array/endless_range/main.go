package main

import "fmt"

func main() {
	data := []int{0, 1, 2}

	//// endless loop
	//// при использование range, указываемое выражение вычисляется
	//// один раз перед началом цикла и копируется во временную переменную
	for range data {
		data = append(data, 10)
		fmt.Println("iteration")
	}

	//// infinity loop
	// for i := 0; i < len(data); i++ {
	// 	data = append(data, 10)
	// 	fmt.Println("iteration")
	// }
}
