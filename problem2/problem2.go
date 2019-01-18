package main

import (
	"fmt"
	"math"
)

var PHI = (1 + SQRT5) / 2
var SQRT5 = math.Sqrt(5)

func fibonacciInverse(f int) int {
	if f < 2 {
		return f
	}
	return int(math.Round(math.Log(float64(f)*SQRT5)/math.Log(PHI))) - 1
}

func fibonacci(n int) int {
	n++
	return int(math.Round(float64(((math.Pow(PHI, float64(n)) - math.Pow(1-PHI, float64(n))) / SQRT5))))
}

func sum() (sum int) {
	max := fibonacciInverse(4000000)
	for i := 1; i <= max; i++ {
		if (i-2)%3 == 0 {
			continue
		}

		sum += fibonacci(i)
	}
	return
}

func main() {
	fmt.Println(sum())
}