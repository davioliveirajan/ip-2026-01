package main

import f"fmt"
import "math"

func main() {
	var soma float64

	for n := 0; n < 64; n++ {
		soma += math.Pow(2, float64(n))
	}
	f.Printf("A soma total é: %.0f\n", soma)
}
