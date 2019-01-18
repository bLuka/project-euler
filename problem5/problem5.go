package main

import "fmt"

func compute() uint {
	const to uint = 20
	var factors = make([]uint, 0, to)
	var result uint = 1

	for i := uint(2); i < to; i++ {
		number := i
		for j := 0; j < len(factors); j++ {
			if number%factors[j] == 0 {
				number /= factors[j]
			}
		}
		factors = append(factors, number)
		result *= number
	}

	return result
}

func main() {
	number := compute()
	fmt.Println(number)
}
