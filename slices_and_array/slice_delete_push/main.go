package main

import "slices"

func main() {
	data := make([]int, 100)

	//// pop front
	data, front := data[1:], data[0] // уменьшается capacity
	_ = front

	//// pop back
	data, back := data[:len(data)-1], data[len(data)-1]
	_ = back

	//// delete via std library
	data = slices.Delete(data, 10, 15) // удалит элементы в диапазоне [10, 15]

	//// manual delete
	data = append(data[:10], data[15:]...)

	//// push front
	data = append([]int{1}, data...)

	//// push back
	data = append(data, 1)
	_ = data
}
