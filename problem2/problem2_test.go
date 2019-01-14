package main

import "testing"

func BenchmarkFibonacciInverse(b *testing.B) {
	test := func(n int) func(*testing.B) {
		return func(b *testing.B) {
			f := fibonacci(n)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				fibonacciInverse(f)
			}
		}
	}

	b.Run("N=10", test(10))
	b.Run("N=20", test(20))
	b.Run("N=30", test(30))
	b.Run("N=50", test(50))
	b.Run("N=70", test(70))
}

func BenchmarkFibonacci(b *testing.B) {
	test := func(n int) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonacci(n)
			}
		}
	}

	b.Run("N=10", test(10))
	b.Run("N=20", test(20))
	b.Run("N=30", test(30))
	b.Run("N=50", test(50))
	b.Run("N=70", test(70))
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum()
	}
}
