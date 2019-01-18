package main

import (
	"fmt"
	"math"
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

func compute() (matchX, matchY uint) {
	var match uint = 100 * 100
	matchX, matchY = 100, 100
	var x, y uint = 999, 999

	for ; x >= 100 && x*y > match; x-- {
		for i := uint(0); (y-i) > x && x*(y-i) > match; i++ {
			if isPalindrome(x * (y - i)) {
				matchX, matchY = x, y-i
				match = matchX * matchY
			}
		}
	}
	return
}

func main() {
	x, y := compute()
	fmt.Printf("%d×%d = %d\n", x, y, x*y)
}
