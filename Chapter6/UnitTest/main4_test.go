package main

import "testing"

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(5)
	}
}
