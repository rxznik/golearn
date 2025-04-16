package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type resultWithError struct {
	user User
	err  error
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

func deactivateUser(wg *sync.WaitGroup, inputCh <-chan User, outputCh chan<- resultWithError) {
	defer wg.Done()

	for user := range inputCh {
		err := user.Deactivate()
		outputCh <- resultWithError{user, err}
	}
}

func DeactivateUsers(users []User, gCount int) ([]User, error) {
	inputCh := make(chan User)
	outputCh := make(chan resultWithError)

	wg := new(sync.WaitGroup)

	go func() {
		defer close(inputCh)

		for idx := range users {
			inputCh <- users[idx]
		}

	}()

	go func() {
		defer close(outputCh)

		for range gCount {
			wg.Add(1)
			go deactivateUser(wg, inputCh, outputCh)
		}

		wg.Wait()
	}()

	output := make([]User, 0, len(users))
	for res := range outputCh {
		if res.err != nil {
			return nil, res.err
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
