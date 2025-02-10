package main

import (
	"bytes"
	"io"
	"testing"
)

func BenchmarkNonAlloc(b *testing.B) {
	n := 100
	for i := 0; i < b.N; i++ {
		ints := []int{}
		for i := 0; i < n; i++ {
			ints = append(ints, i)
		}
	}
}

func BenchmarkSliceAlloc(b *testing.B) {
	n := 100
	for i := 0; i < b.N; i++ {
		ints := make([]int, n)
		for i := 0; i < n; i++ {
			// ints = append(ints, i) 2 allocs/op
			ints[i] = i // 1 alloc/op
		}
	}
}

func BenchmarkWriteToBuffer(b *testing.B) {
	msg := []byte("hello world")
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ { // internal bench stuff, handled by golang

		for i := 0; i < 100; i++ { // our logic
			writeToBuffer(buf, msg)
			buf.Reset()
		}
	}
}
func writeToBuffer(w io.Writer, msg []byte) {
	// buf := new(bytes.Buffer)
	w.Write(msg)
}
