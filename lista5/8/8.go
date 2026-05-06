package main

import f"fmt"
import "math"

func main() {
	var vetor[] float64
	var n int

	f.Printf("Digite 15 números: ")
	for i := 0; i < 15; i++ {
		f.Scan(&n)

		if n < 0 {
			vetor = append(vetor, -1)
		} else {
			vetor = append(vetor, math.Pow(float64(n), 1.0 / 2.0))
		}

		f.Printf("%.2f ", vetor[i])
	}
}
