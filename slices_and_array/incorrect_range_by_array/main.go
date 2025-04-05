package main

import (
	"fmt"
	"unsafe"
)

func main() {
	values := [...]int{100, 200, 300}

	//// Неправильно, также работает и в слайсах
	for idx, val := range values {
		val += 50 // это копия значения, а не ссылка
		//// При печати мы заметим, что адресс у val разный это потому, что,
		//// начиная с Go 1.22, для каждой иттерационной переменной создаётся уникальный экземпляр
		fmt.Println("#1:", unsafe.Pointer(&val), "#2:", unsafe.Pointer(&values[idx]))
	}
	fmt.Println(values) // [100 200 300]

	//// Правильно, также работает и в слайсах
	for idx := range values {
		values[idx] += 50
	}

	fmt.Println(values) // [150 250 350]
}
