package main

import "time"

// check channels in Debug Mode

func main() {
	ch := make(chan uint32, 4)

	ch <- 14
	ch <- 88

	_ = <-ch

	ch <- 11

	_ = <-ch

	go func() {
		for {
			ch <- 52
		}
	}()

	time.Sleep(time.Second)

	return
}
