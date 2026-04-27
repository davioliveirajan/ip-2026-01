package main 

import f"fmt"

func main() {
	l := [] int {1, 4, 6, 8, 10, 11, 15, 19, 20, 15}
	x := 15
	f.Printf("%d\n", buscabinaria(l, x))
}

func buscabinaria(l[] int, x int) int {
	n := len(l)
	e := 0
	d := n - 1
	for e <= d {
		m := (e + d) / 2
		if l[m] == x {
			return m
		}
		if l[m] > x {
			d = m - 1
		}
		if l[m] < x {
			e = m + 1
		}
	}
	return -1
}