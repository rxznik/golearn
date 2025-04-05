# Массивы и срезы в Go ([...]T & []T)

<details>
<summary><b>🏠 Основные понятия</b></summary>

**Массив** - структура данных, которая представляет собой последовательность элементов одного и того же типа, которая имеет фиксированное количество элементов и фиксированный размер элементов. Данные хранятся в памяти подряд.

*Пример создания массива int-элементов в Go размером 5:*
```go
var data [5]int // [0, 0, 0, 0, 0]
```

**Срез (slice)** - структура в Go, имеющая 3 поля:

* **Ссылка на массив** - *array, указатель на первый элемент среза.
* **Длина** - len, количество элементов в срезе.
* **Ёмкость** - cap, максимальное количество элементов в срезе.

*Пример создания нулевого среза int-элементов в Go размером 5 и ёмкостью 10:*

```go
data := make([]int, 5, 10) // [0, 0, 0, 0, 0]
```

</details>


<details>
<summary><b>💫 Аллокация масивов и срезов</b></summary>

**Аллокация массива в кучу** происходит при создании массива размером более 128 КБ, а также при остальных случаях, когда Go аллоцирует данные в кучу (например, при создании массива вне стека, при ситуациях когда ф-ция возвращает открытую ссылку на массив).

*Пример аллокации массива в кучу и на стек:*

```go
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
```

**Аллокация среза в кучу** происходит при создании среза с размером более 64 КБ (но если мы создаём срез от массива не более 128 КБ, то он не будет аллоцироваться в кучу), при реалокации массива в слайсе (т.е. при увеличении ёмкости), а также при остальных случаях, когда Go аллоцирует данные в кучу.

*Пример аллокации среза в кучу и на стек:*

```go
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

	sliceWithHeap := make([]int8, (64<<10)+1) // 64 KB + 1 byte // moved to heap
	_ = sliceWithHeap

	sliceInHeap := allocation() // in heap
	_ = sliceInHeap
}
```

*Пример реаллокации массива в срезе:*

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	slice := make([]int8, 0, 3)
	fmt.Printf("slice: %v\nslice address: %p\n", slice, unsafe.SliceData(slice))

	slice = append(slice, 1, 2, 3)
	fmt.Println("slice full capacity:", slice, "len:", len(slice), "cap:", cap(slice))
	fmt.Println("slice address:", unsafe.SliceData(slice)) // адресс остался неизменен

	/// превосхождение capacity
	slice = append(slice, 4)
	/// адресс массива сильно изменился, произошла аллокация в кучу
	fmt.Printf("slice: %v\nslice address: %p\n", slice, unsafe.SliceData(slice))
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}
```
</details>