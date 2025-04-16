package main

import (
	"fmt"
	"sync"
)

func Tee[T any](inputChannel <-chan T, n int) []chan T {
	outputChannels := make([]chan T, n)
	for i := range n {
		outputChannels[i] = make(chan T)
	}

	go func() {
		for value := range inputChannel {
			for i := range n {
				outputChannels[i] <- value
			}
		}

		for _, channel := range outputChannels {
			close(channel)
		}
	}()

	return outputChannels
}

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)

		for i := range 5 {
			channel <- i
		}
	}()

	const numChannels = 3
	teeChannels := Tee(channel, numChannels)

	var wg sync.WaitGroup

	for i := range teeChannels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range teeChannels[i] {
				fmt.Printf("ch%d: %v\n", i+1, value)
			}
		}()
	}

	wg.Wait()
}
