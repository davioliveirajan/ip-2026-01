package main

import f"fmt"

func main() {
	var s float64

	x := 0

	for i := 100; i > 80; i-- {
		s += float64(i) / float64(fatorial(x))
		x++
	}
	f.Printf("%.2f\n", s)
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}