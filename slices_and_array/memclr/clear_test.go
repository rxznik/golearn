package main

import "testing"

func BenchmarkClearWithFive(b *testing.B) {
	data := make([]int8, 10<<10)
	for i := 0; i < b.N; i++ {
		for idx := range data {
			data[idx] = 5
		}
	}
}

func BenchmarkClearWithZero(b *testing.B) {
	data := make([]int8, 10<<10)
	for i := 0; i < b.N; i++ {
		for idx := range data {
			data[idx] = 0
		}
	}
}
