package main

import f"fmt"

func main() {
	var n [] int
	var x int
	f.Printf("Digite 10 números:")
	for i := 0; i < 10; i++ {
		var num int
		f.Scan(&num)
		n = append(n, num)
	}

	for i, v := range n {
		if n[i] > 50 {
			f.Printf("%d número: %d\n", i + 1, v)
			x++
		}
	}
	if x == 0 {
		f.Printf("Não há números maiores que 50\n")
	}
}