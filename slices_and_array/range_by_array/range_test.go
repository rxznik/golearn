package main

import "testing"

type account struct {
	balance int
}

func BenchmarkWithPointers(b *testing.B) {
	accounts := [...]*account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, a := range accounts {
			a.balance++
		}
	}
}

//// С индексами работает быстрее
func BenchmarkWithIndexes(b *testing.B) {
	accounts := [...]account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for idx := range accounts {
			accounts[idx].balance++
		}
	}
}
