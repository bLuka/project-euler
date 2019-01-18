package main

import (
	"fmt"
	"math"
	"sync"
)

func isPalindrome(n uint) bool {
	switch uint(math.Log10(float64(n))) {
	case 0:
		return true
	case 1:
		return n%10 == n/10
	case 2:
		return n%10 == n/100
	case 3:
		return (n%10 == n/1000) &&
			((n%100-n%10)/10 == n/100-((n/1000)*10))
	case 4:
		return (n%10 == n/10000) &&
			((n%100-n%10)/10 == n/1000-((n/10000)*10))
	case 5:
		return (n%10 == n/100000) &&
			((n%100-n%10)/10 == n/10000-((n/100000)*10)) &&
			((n%1000-n%100)/100 == n/1000-((n/10000)*10))
	}
	return false
}

func compute() (match uint) {
	match = 100 * 100
	const y uint = 999
	var x uint = y
	var matchX uint = 100

	var wg sync.WaitGroup
	for ; x >= matchX; x-- {
		wg.Add(1)
		go (func(x uint) {
			defer wg.Done()
			for i := uint(0); ; i++ {
				yMinusI := y - i
				num := x * yMinusI
				if !(yMinusI > x && num > match) {
					break
				}

				if isPalindrome(num) {
					match = num
					matchX = x
				}
			}
		})(x)
	}
	wg.Wait()
	return
}

func main() {
	match := compute()
	fmt.Printf("%d\n", match)
}
