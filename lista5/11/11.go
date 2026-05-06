package main

import f"fmt"
import "math"

func main() {
	var n, x, somatorio float64
	var vetor[] float64

	for i := 0; i < 100; i++ {
		f.Scan(&n)
		vetor = append(vetor, n)
	}

	x = 99

	for _, v := range vetor {
		somatorio += math.Pow(v - vetor[int(x)], 3)
		x--
	}

	f.Printf("%.6f\n", somatorio)
}
