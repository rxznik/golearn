package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Напишите программу на Go, которая в качестве входных данных получает набор текстов и в 3 потока выводит эти тексты в перевёрнутом виде.
// Входные данные: Каждая строка вывода должна содержать свой порядковый номер (1, 2, 3, ...) и номер потока (1, 2 или 3).
//  Перед завершением работы программа должна вывести "done".

// Тексты: "Hello", "qwerty", "Golang", "platypus", "тест", "level", "generics"

// Пример вывода (обратите внимание на номера линий):

// line 1, thread 1: "olleH"
// line 2, thread 2: "gnaloG"
// line 3, thread 1: "ytrewq"
// line 4, thread 3: "scireneg"
// ...
// done

var Texts = []string{"Hello", "qwerty", "Golang", "platypus", "тест", "level", "generics"}

type Data struct {
	thread   int
	reversed string
}

func main() {
	var (
		wg      sync.WaitGroup
		line    atomic.Int32
		inputCh = make(chan string)
		result  = make([]*Data, len(Texts))
	)

	go func() {
		defer close(inputCh)
		for _, text := range Texts {
			inputCh <- text
		}
	}()

	for i := range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for text := range inputCh {
				result[line.Add(1)-1] = &Data{i + 1, reverse(text)}
				runtime.Gosched() // чтобы 1 горутина всё не выполнила
			}
		}()
	}

	wg.Wait()

	for i, data := range result {
		fmt.Printf("line %d, thread %d: %s\n", i+1, data.thread, data.reversed)
	}
	fmt.Println("done")
}

func reverse(text string) string {
	if len(text) == 0 {
		return ""
	}

	runes := []rune(text)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
