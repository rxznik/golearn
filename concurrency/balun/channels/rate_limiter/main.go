package main

import "time"

type RateLimiter struct {
	leakyBucketCh chan struct{}
	closeCh       chan struct{}
	closeDoneCh   chan struct{}
}

func NewLeakyBucketLimiter(limit int, period time.Duration) RateLimiter {
	limiter := RateLimiter{
		leakyBucketCh: make(chan struct{}, limit),
		closeCh:       make(chan struct{}),
		closeDoneCh:   make(chan struct{}),
	}
	leackInterval := period.Nanoseconds() / int64(limit)
	go limiter.startPeriodicTick(time.Duration(leackInterval))
	return limiter
}

func (rl *RateLimiter) startPeriodicTick(interval time.Duration) {
	timer := time.NewTicker(interval)
	defer func() {
		timer.Stop()
		close(rl.closeDoneCh)
	}()

	for {
		select {
		case <-rl.closeCh:
			return
		default:
		}

		select {
		case <-rl.closeCh:
			return
		case <-timer.C:
			select {
			case <-rl.leakyBucketCh:
			default:
			}
		}
	}
}

func (rl *RateLimiter) Allow() bool {
	select {
	case rl.leakyBucketCh <- struct{}{}:
		return true
	default:
		return false
	}
}

func (rl *RateLimiter) ShutDown() {
	close(rl.closeCh)
	<-rl.closeDoneCh
}

func main() {

}
