package main

import f "fmt"

func main() {
	var h, x float64

	x = 1

	for i := 1; i <= 50; i++ {

		h += x / float64(i)

		x += 2
	}
	f.Printf("H = %.6f\n", h)
}
