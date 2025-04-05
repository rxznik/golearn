package main

import "testing"

const BufferSize = 1024

type ReaderWithSliceArgument struct{}

func (r ReaderWithSliceArgument) Read(p []byte) (int, error) {
	for i := 0; i < BufferSize; i++ {
		p[i] = byte(i)
	}

	return BufferSize, nil
}

type ReaderWithSliceReturn struct{}

func (r ReaderWithSliceReturn) Read(n int) ([]byte, error) {
	p := make([]byte, n)
	for i := 0; i < n; i++ {
		p[i] = byte(i)
	}

	return p, nil
}

func BenchmarkSliceWithArgument(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reader := ReaderWithSliceArgument{}
		_, _ = reader.Read(make([]byte, BufferSize))
	}
}

func BenchmarkSliceWithReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reader := ReaderWithSliceReturn{}
		_, _ = reader.Read(BufferSize)
	}
}
