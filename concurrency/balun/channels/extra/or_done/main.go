package main

import (
	"fmt"
	"time"
)

func OrDone[T any](channel <-chan T, done chan struct{}) <-chan T {
	outputChan := make(chan T)

	go func() {
		defer close(outputChan)
		for {
			select {
			case <-done:
				return
			default:
			}

			select {
			case value, opened := <-channel:
				if !opened {
					return
				}
				outputChan <- value
			case <-done:
				return
			}
		}
	}()

	return outputChan
}

func main() {
	channel := make(chan string)

	go func() {
		for {
			channel <- "test"
			time.Sleep(200 * time.Millisecond)
		}
	}()

	done := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		close(done)
	}()

	for value := range OrDone(channel, done) {
		fmt.Println(value)
	}

}
