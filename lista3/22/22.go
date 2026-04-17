package main

import f"fmt"

func main() {
	var s float64

	for i := 37; i >= 1; i-- {
		n := 1
		s += float64(i) * (float64(i) + 1) / float64(n) 
		n++
	}

	f.Printf("Soma = %.2f\n", s)
}
