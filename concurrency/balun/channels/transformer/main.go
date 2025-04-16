package main

import "fmt"

func Transform[T any](channel <-chan T, action func(T) T) chan T {
	outputChan := make(chan T)

	go func() {
		defer close(outputChan)
		for value := range channel {
			outputChan <- action(value)
		}
	}()

	return outputChan
}

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := range 5 {
			channel <- i
		}
	}()

	mul := func(value int) int {
		return value * value
	}

	for value := range Transform(channel, mul) {
		fmt.Println(value)
	}
}
