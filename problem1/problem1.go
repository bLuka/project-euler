package main

import "fmt"

const MAX = 1000

func compute() uint {
	var max uint = MAX - 1

	var max_seq3 uint = max - max%3
	var max_seq5 uint = max - max%5
	var max_seq15 uint = max - max%15

	var nb_items_seq3 uint = max_seq3 / 3
	var nb_items_seq5 uint = max_seq5 / 5
	var nb_items_seq15 uint = max_seq15 / 15

	var mod2th_max_seq3 uint = max_seq3 - (nb_items_seq3%2)*3
	var mod2th_max_seq5 uint = max_seq5 - (nb_items_seq5%2)*5
	var mod2th_max_seq15 uint = max_seq15 - (nb_items_seq15%2)*15

	var median_seq3 uint = nb_items_seq3 / 2
	var median_seq5 uint = nb_items_seq5 / 2
	var median_seq15 uint = nb_items_seq15 / 2

	return (mod2th_max_seq3+3)*median_seq3 + (nb_items_seq3%2)*max_seq3 + // Sum of multiples of 3…
		(mod2th_max_seq5+5)*median_seq5 + (nb_items_seq5%2)*max_seq5 - // …plus sum of multiples of 5…
		((mod2th_max_seq15+15)*median_seq15 + (nb_items_seq15%2)*max_seq15) // …minus sum of multiples of 15
}

func main() {
	fmt.Println(compute())
}
