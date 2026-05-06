package main

import f"fmt"

func main() {
	var vetor[] int
	var n int

	for i := 0; i < 20; i++ {
		f.Scan(&n)
		vetor = append(vetor, n)
	}

	frequencia := make(map[int]int)

	for _, v := range vetor {
		frequencia[v]++
	}

	for vetor, freq := range frequencia {
		f.Printf("Número %d: %d vezes\n", vetor, freq)
	}
}
