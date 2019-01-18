package main

import (
	"math"
)

func isPalindrome(n int) bool {
	switch int(math.Log10(float64(n))) {
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

func compute() {
}

func main() {
	compute()
}
