package main

import (
	"fmt"
	"sync"
)

func Bridge[T any](channel chan chan T) <-chan T {
	outputChan := make(chan T)

	go func() {
		var wg sync.WaitGroup
		for ch := range channel {
			wg.Add(1)
			go func(ch chan T) {
				defer wg.Done()
				for value := range ch {
					outputChan <- value
				}
			}(ch)
		}

		go func() {
			wg.Wait()
			close(outputChan)
		}()
	}()

	return outputChan
}

func main() {
	channelChannel := make(chan chan string)

	go func() {
		channel1 := make(chan string, 3)
		for range 3 {
			channel1 <- "channel-1"
		}

		close(channel1)

		channel2 := make(chan string, 3)
		for range 3 {
			channel2 <- "channel-2"
		}

		close(channel2)
		channelChannel <- channel1
		channelChannel <- channel2

		close(channelChannel)
	}()

	for value := range Bridge(channelChannel) {
		fmt.Println(value)
	}

}
