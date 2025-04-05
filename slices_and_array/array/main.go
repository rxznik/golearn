package main

import "fmt"

func main() {
	//// создание массива, размером может быть только константа
	//// нельзя создавать с помощью make и использовать append
	const SIZE = 5

	var data1 [SIZE]int
	fmt.Println(data1) // [0 0 0 0 0]

	var data2 [2][SIZE]int
	fmt.Println(data2) // [[0 0 0 0 0] [0 0 0 0 0]]

	data3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(data3) // [1 2 3 4 5]

	data4 := [SIZE]int{1, 2, 3}
	fmt.Println(data4) // [1 2 3 0 0]

	data5 := [SIZE]int{3: 4}
	fmt.Println(data5) // [0 0 0 4 0]

	data6 := [SIZE]int{3: 4, 5, 1: 2}
	fmt.Println(data6) // [0 2 0 4 5]

	//// получение значения и ссылки по индексу
	val := data6[3]
	valPtr := &data6[3]
	fmt.Printf("value: %d, pointer: %p\n", val, valPtr) // value: 4, pointer: <smth address>

	//// изменение значения по индексу
	data6[0]++
	data6[2] = 3
	fmt.Println(data6) // [1 2 3 4 5]

	//// получение длины и ёмкости массива
	fmt.Printf("len: %d, cap: %d\n", len(data6), cap(data6)) // len: 5, cap: 5

	//// получение среза
	s := data6[3:5]
	fmt.Println(s) // [4 5]

	//// сравнение массивов
	fmt.Printf("data3 == data6: %t\n", data3 == data6) // data3 == data6: true
	fmt.Printf("data1 == data6: %t\n", data1 == data6) // data1 == data6: false
}
