package main

import f"fmt"
import "math"

func main() {
	var s float64

	x := 15

	for i := 0; i <= 14 && x > 1; i += 2 {
		s += (math.Pow(2, float64(i)) / math.Pow(float64(x), float64(2))) - (math.Pow(2, float64(i+1)) / math.Pow(float64(x-1), float64(2)))
		x -= 2
	}

	s += math.Pow(2, float64(14))
	
	f.Printf("%.2f\n", s)
}