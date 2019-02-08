package main

import (
	"fmt"
)

func is6Digits(n uint) bool {
	return n >= 100000
}

func isPalindrome(n uint) bool {
	if is6Digits(n) {
		return (n%10 == n/100000) &&
			((n%100-n%10)/10 == n/10000-((n/100000)*10)) &&
			((n%1000-n%100)/100 == n/1000-((n/10000)*10))
	}
	return (n%10 == n/10000) &&
		((n%100-n%10)/10 == n/1000-((n/10000)*10))
}

func compute() (match uint) {
	match = 100 * 100
	const y uint = 999
	var x uint = y

	for ; x*y >= match; x-- {
		for i := uint(0); ; i++ {
			yMinusI := y - i
			num := x * yMinusI
			if !(yMinusI >= x && num > match) {
				break
			}

			if isPalindrome(num) {
				match = num
			}
		}
	}
	return
}

func main() {
	match := compute()
	fmt.Printf("%d\n", match)
}
