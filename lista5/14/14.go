package main

import f"fmt"

func main() {
	var vetor1[] int 
	var vetor2[] int
	var result[] int
	var n int

	f.Printf("Digite 10 números: ")
	for i := 0; i < 10; i++ {
		f.Scan(&n)
		vetor1 = append(vetor1, n)
	}

	f.Printf("Digite 10 números: ")
	for i := 0; i < 10; i++ {
		f.Scan(&n)
		vetor2 = append(vetor2, n)
	}

	x := 0
	y := 0

	for i := 0; i < 20; i++ {
		if i == 0 || i % 2 == 0 {
			result = append(result, vetor1[x])
			x++
		} else {
			result = append(result, vetor2[y])
			y++
		}
	}
	f.Println(result)
}

