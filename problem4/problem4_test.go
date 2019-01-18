package main

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

var testingMap = map[uint]bool{
	1: true,
	5: true,

	11: true,
	55: true,
	99: true,

	101: true,
	181: true,
	929: true,

	1221: true,
	2002: true,
	6556: true,
	9229: true,

	19291: true,
	58285: true,
	99099: true,
	99999: true,

	123321: true,
	189981: true,
	989989: true,
	999999: true,

	10:     false,
	100:    false,
	1000:   false,
	10000:  false,
	100000: false,

	12:     false,
	110:    false,
	1212:   false,
	1222:   false,
	11221:  false,
	211221: false,
}

func TestIsPalindrome(t *testing.T) {
	test := func(n uint, expected bool) func(*testing.T) {
		return func(t *testing.T) {
			if isPalindrome(n) != expected {
				t.Fail()
			}
		}
	}

	for n, expected := range testingMap {
		t.Run(fmt.Sprintf("N=%d", n), test(n, expected))
	}
}

func TestCompute(t *testing.T) {
	if compute() != 906609 {
		t.Fatal("compute() returned", compute())
	}
}

func BenchmarkCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compute()
	}
}

func BenchmarkLog10(b *testing.B) {
	bench := func(n uint) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				math.Log10(float64(n))
			}
		}
	}

	keys := make([]uint, 0, len(testingMap))
	for k := range testingMap {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, n := range keys {
		b.Run(fmt.Sprintf("N=%d", n), bench(n))
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	bench := func(n uint) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPalindrome(n)
			}
		}
	}

	validKeys := make([]uint, 0, len(testingMap))
	invalidKeys := make([]uint, 0, len(testingMap))
	for k, expected := range testingMap {
		if expected {
			validKeys = append(validKeys, k)
		} else {
			invalidKeys = append(invalidKeys, k)
		}
	}
	sort.Slice(validKeys, func(i, j int) bool { return validKeys[i] < validKeys[j] })
	sort.Slice(invalidKeys, func(i, j int) bool { return invalidKeys[i] < invalidKeys[j] })

	for _, n := range validKeys {
		b.Run(fmt.Sprintf("%t/log10N=%d/N=%d", testingMap[n], uint(math.Log10(float64(n))), n), bench(n))
	}
	for _, n := range invalidKeys {
		b.Run(fmt.Sprintf("%t/log10N=%d/N=%d", testingMap[n], uint(math.Log10(float64(n))), n), bench(n))
	}
}
