package main

import (
	"testing"
	"unsafe"
)

func Convert(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	return unsafe.String(unsafe.SliceData(data), len(data))
}

var Result string

func BenchmarkConvertion(b *testing.B) {
	slice := []byte("Hello world!")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Result = string(slice)
	}
}

func BenchmarkUnsafeConvertion(b *testing.B) {
	slice := []byte("Hello world!")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Result = Convert(slice)
	}
}
