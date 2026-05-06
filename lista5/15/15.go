package main

import f"fmt"

func main() {
	var vetor[] int
	var n int
	for i := 0; i < 30; i++ {
		f.Scan(&n)
		vetor = append(vetor, n)
	}

	x := 1

	for i, v := range vetor {
		if x % 2 == 0 {
			vetor[i] = 2 * v
			x++
		} else if x % 2 == 1 {
			vetor[i] = 3 * v
			x++
		}
	}

	f.Println(vetor)
}
