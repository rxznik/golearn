package main

import "fmt"

func Filter[T any](channel <-chan T, predicate func(T) bool) <-chan T {
	outputChan := make(chan T)

	go func() {
		defer close(outputChan)
		for value := range channel {
			if predicate(value) {
				outputChan <- value
			}
		}
	}()

	return outputChan
}

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := range 10 {
			channel <- i
		}
	}()

	checkEven := func(value int) bool {
		if value%2 == 0 {
			return true
		}
		return false
	}

	for value := range Filter(channel, checkEven) {
		fmt.Println(value)
	}
}
