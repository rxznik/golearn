package main

import (
	"fmt"
	"sync"
)

func split[T any](channel <-chan T, n int) []chan T {
	outputChannels := make([]chan T, n)

	for i := range n {
		outputChannels[i] = make(chan T)
	}

	go func() {
		defer func() {
			for i := range outputChannels {
				close(outputChannels[i])
			}
		}()

		idx := 0
		for value := range channel {
			outputChannels[idx] <- value
			idx = (idx + 1) % n
		}
	}()

	return outputChannels
}

func parse(channel <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		defer close(outputChannel)
		for value := range channel {
			outputChannel <- fmt.Sprintf("parsed - %s", value)
		}
	}()

	return outputChannel
}

func send(channel <-chan string, n int) <-chan string {
	var wg sync.WaitGroup
	splittedChannels := split(channel, n)
	outputChannel := make(chan string)

	wg.Add(n)
	for i := range splittedChannels {
		go func(idx int) {
			defer wg.Done()
			for value := range splittedChannels[idx] {
				outputChannel <- fmt.Sprintf("sent - %s", value)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(outputChannel)
	}()

	return outputChannel
}

func main() {
	channel := make(chan string)

	go func() {
		defer close(channel)
		for i := range 25 {
			channel <- fmt.Sprintf("user_%d", i)
		}
	}()

	for value := range send(parse(channel), 5) {
		fmt.Println(value)
	}
}
