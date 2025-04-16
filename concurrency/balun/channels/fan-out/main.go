package main

import (
	"fmt"
	"sync"
)

func SplitChannel[T any](inputChannel <-chan T, n int) []chan T {
	outputChannels := make([]chan T, n)
	for i := range n {
		outputChannels[i] = make(chan T)
	}

	go func() {
		idx := 0
		for value := range inputChannel {
			outputChannels[idx] <- value
			idx = (idx + 1) % n
		}

		for _, ch := range outputChannels {
			close(ch)
		}
	}()

	return outputChannels
}

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := range 100 {
			channel <- i
		}
	}()

	const channelsCount = 25

	channels := SplitChannel(channel, channelsCount)

	var wg sync.WaitGroup

	for i := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range channels[i] {
				fmt.Printf("ch%d: %v\n", i+1, value)
			}
		}()
	}

	wg.Wait()
}
