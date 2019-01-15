package main

import "testing"

func BenchmarkCompute(b *testing.B) {
	test := func(N int) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				compute(N)
			}
		}
	}

	b.Run("Compute(10)", test(10))
	b.Run("Compute(100)", test(100))
	b.Run("Compute(1000)", test(1000))
	b.Run("Compute(13195)", test(13195))
	b.Run("Compute(600851475143)", test(600851475143))
}
