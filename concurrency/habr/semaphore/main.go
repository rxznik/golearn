package main

import (
	"fmt"
	"log"
	"time"
)

type Semaphore struct {
	C chan struct{}
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{make(chan struct{}, n)}
}

func (s *Semaphore) Acquire() {
	s.C <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.C
}

type User struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

func (u *User) Activate() error {
	time.Sleep(time.Millisecond * 200)
	u.Active = true
	return nil
}

func (u *User) Deactivate() error {
	time.Sleep(time.Millisecond * 200)
	u.Active = false
	return nil
}

type ResultWithError struct {
	user User
	err  error
}

func DeactivateUsers(users []User, gCount int) ([]User, error) {
	sem := NewSemaphore(gCount)

	outputCh := make(chan ResultWithError)
	doneCh := make(chan struct{})

	output := make([]User, 0, len(users))
	for _, user := range users {
		go func(user User) {
			sem.Acquire()
			defer sem.Release()
			err := user.Deactivate()

			select {
			case outputCh <- ResultWithError{user, err}:
			case <-doneCh:
				return
			}
		}(user)
	}

	for range len(users) {
		res := <-outputCh
		if res.err != nil {
			close(doneCh)
			return nil, fmt.Errorf("DeactivateUsers: %v", res.err)
		}

		output = append(output, res.user)
	}
	return output, nil
}

func main() {
	users := []User{
		{Name: "Alice", Active: true},
		{Name: "Joker", Active: true},
		{Name: "Bruce", Active: true},
		{Name: "Tayler", Active: true},
		{Name: "Den", Active: true},
		{Name: "Kate", Active: true},
	}

	output, err := DeactivateUsers(users, 5)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(output)
}
