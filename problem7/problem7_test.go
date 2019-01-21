package main

import "testing"

var numbers = map[uint]bool{
	1:  true,
	2:  true,
	3:  true,
	5:  true,
	7:  true,
	11: true,
	13: true,

	4:  false,
	6:  false,
	8:  false,
	9:  false,
	10: false,
}

func TestIsPrime(t *testing.T) {
	for n, result := range numbers {
		if isPrime(n) != result {
			t.Error("isPrime() of", n, "unexpectedly returned", !result)
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	test := func(number uint) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPrime(number)
			}
		}
	}

	b.Run("isPrime(10)", test(10))
	b.Run("isPrime(100)", test(100))
	b.Run("isPrime(1000)", test(1000))
	b.Run("isPrime(10000)", test(10000))
	b.Run("isPrime(100000)", test(100000))
	b.Run("isPrime(1000000)", test(1000000))
}

func BenchmarkCompute(b *testing.B) {
	test := func(number uint) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				compute(number)
			}
		}
	}

	b.Run("compute(10)", test(10))
	b.Run("compute(100)", test(100))
	b.Run("compute(1000)", test(1000))
	b.Run("compute(10000)", test(10000))
}
