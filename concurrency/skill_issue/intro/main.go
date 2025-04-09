package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func randomWait() time.Duration {
	workTime := time.Duration(rand.Intn(5)+1) * time.Second
	time.Sleep(workTime)
	return workTime
}

const OperationsCount = 100

func solutionViaChannel() {
	var (
		executionTime time.Duration
		totalTime     time.Duration
		wg            sync.WaitGroup
		ch            = make(chan time.Duration)
		ctx, cancel   = context.WithCancel(context.Background())
	)

	defer cancel()

	start := time.Now()

	for range OperationsCount {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- randomWait()
		}(&wg)
	}

	go func() {
		wg.Wait()
		close(ch)
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			executionTime = time.Since(start)
			fmt.Println("execution time:", executionTime)
			fmt.Println("total time:", totalTime)
			return
		case workTime := <-ch:
			totalTime += workTime
		}
	}
}

func solutionViaBufferedChannel() {
	var (
		executionTime time.Duration
		totalTime     time.Duration
		ch            = make(chan time.Duration, OperationsCount)
	)

	start := time.Now()

	for range OperationsCount {
		go func() {
			ch <- randomWait()
		}()
	}

	for range OperationsCount {
		workTime := <-ch
		totalTime += workTime
	}

	executionTime = time.Since(start)

	fmt.Println("execution time:", executionTime)
	fmt.Println("total time:", totalTime)
}

func solutionViaAtomic() {
	var (
		executionTime time.Duration
		totalTime     time.Duration
		wg            sync.WaitGroup
	)

	start := time.Now()

	for range OperationsCount {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			workTime := randomWait()
			atomic.AddInt64((*int64)(&totalTime), int64(workTime))
		}(&wg)
	}

	wg.Wait()

	executionTime = time.Since(start)

	fmt.Println("execution time:", executionTime)
	fmt.Println("total time:", totalTime)
}

func solutionViaMutex() {
	var (
		executionTime time.Duration
		totalTime     time.Duration
		wg            sync.WaitGroup
		mu            sync.Mutex
	)

	start := time.Now()

	for range OperationsCount {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			workTime := randomWait()
			mu.Lock()
			totalTime += workTime
			mu.Unlock()
		}(&wg)
	}

	wg.Wait()

	executionTime = time.Since(start)

	fmt.Println("execution time:", executionTime)
	fmt.Println("total time:", totalTime)
}

func main() {
	fmt.Println("\x1b[33msolution via channel:\x1b[0m")
	solutionViaChannel()

	fmt.Println("\x1b[33msolution via buffered channel:\x1b[0m")
	solutionViaBufferedChannel()

	fmt.Println("\x1b[33msolution via atomic:\x1b[0m")
	solutionViaAtomic()

	fmt.Println("\x1b[33msolution via mutex:\x1b[0m")
	solutionViaMutex()
}
