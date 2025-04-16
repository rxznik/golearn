package main

import "fmt"

func GenerateWithChannel(start, end int) chan int {
	outputCh := make(chan int)

	go func() {
		defer close(outputCh)
		for i := start; i <= end; i++ {
			outputCh <- i
		}
	}()

	return outputCh
}

func main() {
	for num := range GenerateWithChannel(1400, 1488) {
		fmt.Println(num)
	}
}
