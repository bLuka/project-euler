package main

import (
	"fmt"
	"math/rand"
	"testing"
)

const maxSumWithoutOverflow = 32768
const benchmarkQuantity = 8

func TestCompute(t *testing.T) {
	var maxRuns = 10000

	if testing.Short() {
		maxRuns = 100
	}

	for i := 0; i < maxRuns; i++ {
		random := uint(rand.Intn(maxSumWithoutOverflow))
		t.Run(fmt.Sprintf("compute(%d)", random), func(t *testing.T) {
			expected := longCompute(random)
			if result := compute(random); result != expected {
				t.Fatal("compute() failed, returning", result, "; expecting", expected)
			}
		})
	}
}

func BenchmarkLongCompute(b *testing.B) {
	run := func(maximum uint) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				longCompute(maximum)
			}
		}
	}

	for maximum := uint(0); maximum <= maxSumWithoutOverflow; maximum += maxSumWithoutOverflow / benchmarkQuantity {
		b.Run(fmt.Sprintf("longCompute(%d)", maximum), run(maximum))
	}
}

func BenchmarkCompute(b *testing.B) {
	run := func(maximum uint) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				compute(maximum)
			}
		}
	}

	for maximum := uint(0); maximum <= maxSumWithoutOverflow; maximum += maxSumWithoutOverflow / benchmarkQuantity {
		b.Run(fmt.Sprintf("compute(%d)", maximum), run(maximum))
	}
}
