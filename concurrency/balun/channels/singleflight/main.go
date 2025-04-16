package main

import (
	"fmt"
	"sync"
	"time"
)

type call struct {
	err  error
	val  interface{}
	done chan struct{}
}

type SingleFlight struct {
	mutex sync.Mutex
	calls map[string]*call
}

func NewSingleFlight() *SingleFlight {
	return &SingleFlight{calls: make(map[string]*call)}
}

func (sf *SingleFlight) Do(key string, action func() (interface{}, error)) (interface{}, error) {
	sf.mutex.Lock()

	if call, found := sf.calls[key]; found {
		sf.mutex.Unlock()
		return sf.wait(call)
	}

	call := &call{done: make(chan struct{})}
	sf.calls[key] = call
	sf.mutex.Unlock()

	go func() {
		defer func() {
			sf.mutex.Lock()
			close(call.done)
			delete(sf.calls, key)
			sf.mutex.Unlock()
		}()

		call.val, call.err = action()
	}()

	return sf.wait(call)
}

func (sf *SingleFlight) wait(call *call) (interface{}, error) {
	<-call.done
	return call.val, call.err
}

const numInFlightRequests = 5

func main() {
	var wg sync.WaitGroup
	wg.Add(numInFlightRequests)

	singleFlight := NewSingleFlight()
	const key = "same_key"

	for i := range numInFlightRequests {
		go func(idx int) {
			defer wg.Done()
			value, err := singleFlight.Do(key, func() (interface{}, error) {
				fmt.Println("Single Flight")
				time.Sleep(5 * time.Second)
				return "Joker", nil
			})

			fmt.Printf("response[%d]: %v | %v\n", idx, value, err)
		}(i)
	}
	wg.Wait()
}
