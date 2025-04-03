package main

import "fmt"

func main() {
	//// Нельзя создать nil map, т.к. header map не определен
	// var m map[string]int
	// m["a"] = 1 // panic
	// fmt.Println(m)

	// Правильное создание пустой map.
	// Задаём размер (допустим ожидаем максимум 12 пар)
	// для избежания лишних эвакуаций
	m := make(map[string]int, 12)
	// Вставка значения
	m["apple"] = 1
	fmt.Println(m)

	/// Нельзя ссылаться на значение map по ключу, т.к. при эвакуации
	/// адресс пары ключ-значение может поменяться
	// apple := &m["apple"] // invalid operation

	// Можно получить значение по ключу (1 способ)
	v, ok := m["banana"]
	if !ok {
		fmt.Println("banana not found")
	}

	m["banana"] = 2
	m["orange"] = 3

	//// Можно получить значение по ключу (2 способ)
	v = m["banana"]
	fmt.Println(v)

	//// Примеры изменения значения по ключу
	m["apple"]++
	m["banana"]--
	m["orange"] = m["apple"] + 10
	fmt.Println(m)

	//// Удаление пары ключ-значение
	delete(m, "apple")
	fmt.Println(m)
	//// Удаление всех пар
	clear(m)

	//// При каждом выводе с помощью range, пары ключ-значение будет перемешиваться,
	//// это происходит из-за отсутствия порядка в map,
	//// при итеррации по map начало ( 1-й бакет/группа) определяется с помощью рандома
	m = map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
		"peach":  4,
		"grape":  5,
		"kiwi":   6,
	}

	for i := 0; i < 10; i++ {
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("Hello, user_%d!", i)
		}()
	}

	//// Но если мы выводим map c помощью fmt.Println,
	//// то порядок будет отсортирован по возрастанию ключей
	fmt.Println(m)
}
