package main

import (
	"fmt"
	"time"
)

func or[T any](channels ...<-chan T) <-chan T {
	if len(channels) == 0 {
		return nil
	}

	if len(channels) == 1 {
		return channels[0]
	}

	doneCh := make(chan T)

	go func() {
		defer close(doneCh)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or[T](channels[3:]...):
			}
		}
	}()

	return doneCh
}

func main() {
	start := time.Now()
	<-or(
		time.After(12*time.Second),
		time.After(3*time.Second),
		time.After(45*time.Second),
		time.After(7*time.Second),
		time.After(1*time.Second),
	)
	fmt.Printf("Called after: %s\n", time.Since(start))
}
