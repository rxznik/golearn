package main

import (
	"fmt"
	"runtime"
	"time"
)

// go build; GODEBUG=schedtrace=1000 ./scheduler

func main() {
	fmt.Println(runtime.NumCPU())      // кол-во ядер
	fmt.Println(runtime.GOMAXPROCS(0)) // кол-во P по умолчанию
	runtime.GOMAXPROCS(2)              // установка кол-ва P
	fmt.Println(runtime.GOMAXPROCS(0)) // кол-во P

	// LRQ size = 256

	for range 256 {
		go func() {
			for {
				for i := 0; i < 10000000; i++ {
					_ = i * i
				}
			}
		}()
	}

	fmt.Println(runtime.NumGoroutine()) // кол-во горутин

	time.Sleep(time.Minute)
}
