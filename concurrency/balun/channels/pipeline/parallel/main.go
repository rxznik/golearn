package main

import (
	"fmt"
	"sync"
)

func parse(channel <-chan string) <-chan string {
	outputChan := make(chan string)

	go func() {
		defer close(outputChan)
		for value := range channel {
			outputChan <- fmt.Sprintf("parsed - %s", value)
		}
	}()

	return outputChan
}

func send(channel <-chan string, n int) <-chan string {
	var wg sync.WaitGroup

	outputChan := make(chan string)
	wg.Add(n)
	for range n {
		go func() {
			defer wg.Done()
			for value := range channel {
				outputChan <- fmt.Sprintf("sent - %s", value)
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
	channel := make(chan string)
	go func() {
		defer close(channel)
		for i := range 10 {
			channel <- fmt.Sprintf("user_%d", i)
		}
	}()
	for value := range send(parse(channel), 5) {
		fmt.Println(value)
	}
}
