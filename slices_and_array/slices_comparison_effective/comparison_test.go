package main

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

// / интерфейс Comparable определяет типы, которые можно сравнивать
// / разумеется их больше, но в данном случае мы ограничиваемся
type Comparable interface {
	string | int | float64 | bool | byte
}

func equal[T Comparable](lhs, rhs []T) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}

	return true
}

func BenchmarkWithEqualFunction(b *testing.B) {
	lhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = equal(lhs, rhs)
	}
}

func BenchmarkWithSlicesEqual(b *testing.B) {
	lhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = slices.Equal(lhs, rhs)
	}
}

func BenchmarkWithReflectEqual(b *testing.B) {
	lhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = reflect.DeepEqual(lhs, rhs)
	}
}

func BenchmarkWithSprint(b *testing.B) {
	lhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rhs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(lhs) == fmt.Sprint(rhs)
	}
}
