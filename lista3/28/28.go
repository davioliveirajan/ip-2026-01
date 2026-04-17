package main

import f"fmt"
import "math"

func main() {
	var s, pi float64

	x := 1

	for i := 0; i <= 50; i++ {
		s += 1 / math.Pow(float64(x), 3) - 1 / math.Pow(float64(x+2), 3)
		x += 4
	}

	pi = math.Pow(s * 32, 1.0/3.0)

	f.Printf("%.6f\n", pi)
}