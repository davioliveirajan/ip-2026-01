package main

import f"fmt"

func main() {
	l := []int {5, 2, 7, 9, 10, 56, 98, 1, 4, 6}
	x := 10
	f.Printf("%d\n", buscasequencial(l, x))
}

func buscasequencial(l[] int, x int)int {
	n := len(l)
	for i := 0; i <= n - 1; i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
}