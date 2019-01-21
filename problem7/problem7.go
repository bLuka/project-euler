package main

import (
	"fmt"
	"sort"
)

func isPrime(number uint) bool {
	for i := uint(2); i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func findPrime(start uint, continueTo *uint) func(chan<- uint, chan<- uint) {
	return func(primes chan<- uint, finished chan<- uint) {
		for cursor := start; *continueTo == 0 || cursor <= *continueTo; cursor += 10 {
			if isPrime(cursor) {
				primes <- cursor
			}
		}
		primes <- 0
	}
}

// compute finds the nth prime number
func compute(n uint) uint {
	var primes = make([]uint, 0, n*2)
	var continueTo = uint(0)

	var primesChannel = make(chan uint, n*2)
	defer close(primesChannel)

	var finished = make(chan uint, 4)
	defer close(finished)

	go findPrime(11, &continueTo)(primesChannel, finished)
	go findPrime(13, &continueTo)(primesChannel, finished)
	go findPrime(17, &continueTo)(primesChannel, finished)
	go findPrime(19, &continueTo)(primesChannel, finished)

	var maxPrime = uint(7)
	primes = append(primes, 2, 3, 5, 7)

	for running := uint(4); running != 0; {
		select {
		case prime := <-primesChannel:
			if prime == 0 {
				running--
				continue
			}
			primes = append(primes, prime)

			if prime > maxPrime {
				maxPrime = prime
			}
			if len(primes) == int(n+1) {
				continueTo = maxPrime
			}
		}
	}

	sort.Slice(primes, func(i, j int) bool { return primes[i] < primes[j] })

	return primes[n]
}

func main() {
	fmt.Println(compute(10000))
}
