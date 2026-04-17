package main

import f"fmt"

func main () {
	var n int
	var serie, i, x float64
	f.Printf("Digite a quantidade de termos: ")
	f.Scan(&n)

	i = 1000
	x = 1

	for n > 0 {
		serie += (i / x) - ((i-3) / (x + 1))
		i -= 6
		x += 2
		n--
	}

	f.Printf("%.2f\n", serie)
}
