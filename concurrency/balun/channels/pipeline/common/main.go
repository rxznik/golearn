package main

import "fmt"

func generate[T any](values ...T) <-chan T {
	outputChan := make(chan T)

	go func() {
		defer close(outputChan)
		for _, value := range values {
			outputChan <- value
		}
	}()

	return outputChan
}

func process[T any](channel <-chan T, action func(T) T) <-chan T {
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
	values := []int{1, 2, 3, 4, 5}
	square := func(num int) int {
		return num * num
	}
	for value := range process(generate(values...), square) {
		fmt.Println(value)
	}
}
