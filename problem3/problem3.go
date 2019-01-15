package main

import (
	"fmt"
)

func compute(input int) int {
	for i := 2; i < input; i++ {
		if input%i == 0 {
			input /= i
		}
	}
	return input
}

func main() {
	fmt.Println(compute(600851475143))
}
