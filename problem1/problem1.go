package main

import "fmt"

const MAX = 1000

func compute() int {
	max := MAX - 1

	max_seq3 := max - max%3
	max_seq5 := max - max%5
	max_seq15 := max - max%15

	nb_items_seq3 := max_seq3 / 3
	nb_items_seq5 := max_seq5 / 5
	nb_items_seq15 := max_seq15 / 15

	mod2th_max_seq3 := max_seq3 - (nb_items_seq3%2)*3
	mod2th_max_seq5 := max_seq5 - (nb_items_seq5%2)*5
	mod2th_max_seq15 := max_seq15 - (nb_items_seq15%2)*15

	median_seq3 := nb_items_seq3 / 2
	median_seq5 := nb_items_seq5 / 2
	median_seq15 := nb_items_seq15 / 2

	return (mod2th_max_seq3+3)*median_seq3 + (nb_items_seq3%2)*max_seq3 + // Sum of multiples of 3…
		(mod2th_max_seq5+5)*median_seq5 + (nb_items_seq5%2)*max_seq5 - // …plus sum of multiples of 5…
		((mod2th_max_seq15+15)*median_seq15 + (nb_items_seq15%2)*max_seq15) // …minus sum of multiples of 15
}

func main() {
	fmt.Println(compute())
}
