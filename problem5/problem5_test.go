package main

import "testing"

func TestCompute(t *testing.T) {
	if result := compute(); result != 232792560 {
		t.Fatal("compute() returned", result)
	}
}

func BenchmarkCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compute()
	}
}
