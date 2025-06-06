package main

import (
	"strings"
	"testing"
	"unsafe"
)

func makeDirty(size int) []byte {
	var sb strings.Builder
	sb.Grow(size)

	pointer := unsafe.StringData(sb.String())
	return unsafe.Slice(pointer, size)
}

var Result []byte

func BenchmarkMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = make([]byte, 0, 10<<20) // 10 MB
	}
}

func BenchmarkMakeDirty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Result = makeDirty(10 << 20) // 10 MB
	}
}
