package main

import (
	"fmt"
	"sync"
)

func MergeChannels[T any](channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	wg.Add(len(channels))

	outputChan := make(chan T)
	for _, channel := range channels {
		go func() {
			defer wg.Done()

			for value := range channel {
				outputChan <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	return outputChan
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
		}()

		for i := 0; i < 100; i += 3 {
			ch1 <- i
			ch2 <- i + 1
			ch3 <- i + 2
		}
	}()

	for value := range MergeChannels(ch1, ch2, ch3) {
		fmt.Println(value)
	}
}
