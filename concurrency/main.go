package main

import (
	"fmt"
	"sync"
)

func work(wg *sync.WaitGroup, inputCh <-chan int, outputCh chan<- int) {
	defer wg.Done()
	// пытаемся читать данные, пока канал не закрыт
	for value := range inputCh {
		outputCh <- value * value
	}
}

func main() {
	resultCh := make(chan int)
	inputCh := make(chan int)

	var wg sync.WaitGroup

	// асинхронно записываем данные в inputCh
	go func() {
		// закрываем канал, после того, как все записали, для предотвращения deadlock
		// deadlock может возникнуть из-за того что все receiver-ы будут вечно ждать данные
		defer close(inputCh)
		for i := range 10 {
			inputCh <- i
		}
	}()

	for range 5 {
		wg.Add(1)
		// асинхронно читаем из inputCh и записываем результаты в resultCh
		go work(&wg, inputCh, resultCh)
	}

	go func() {
		wg.Wait()       // ждем пока все горутины завершатся
		close(resultCh) // после этого закрываем канал
	}()

	// при отсутствии данных горутина main заблокируется
	// при закрытии канала чтение завершается
	// в остальных случаях будет читать данные, поступающие от work()
	for value := range resultCh {
		fmt.Println(value)
	}
}
