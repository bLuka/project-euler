package main

import "fmt"

func compute(maximum uint) uint {
	return (3*maximum*maximum*maximum*maximum +
		2*maximum*maximum*maximum -
		3*maximum*maximum -
		2*maximum) / 12
}

func longCompute(maximum uint) uint {
	sumSquares := uint(1)
	sumSquared := uint(1)
	for i := uint(2); i <= maximum; i++ {
		sumSquares += i * i
		sumSquared += i
	}

	sumSquared *= sumSquared
	return sumSquared - sumSquares
}

func main() {
	fmt.Println(compute(100))
}
