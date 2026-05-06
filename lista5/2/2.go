package main

import f"fmt"

func main() {
	var vet1[] int
	var vet2[] int
	var vetresult1 [] int
	var vetresult2 [] int
	var n1, n2, soma int

	f.Printf("Primeiro vetor: ")
	for i := 0; i < 10; i++ {
		f.Scan(&n1)
		vet1 = append(vet1, n1)
	}

	f.Printf("Segundo vetor: ")
	for i := 0; i < 5; i++ {
		f.Scan(&n2)
		vet2 = append(vet2, n2)
	}

	for _, v := range vet2 {
		soma += v
	}

	for _, v := range vet1 {
		if v % 2 == 0 {
			vetresult1 = append(vetresult1, v+soma)
		} else {
			vetresult2 = append(vetresult2, v+soma)
		}
	}

	f.Printf("Primeiro vetor resultante: [ ")
	for _, v := range vetresult1 {
		f.Printf("%d ", v)
	}
	f.Printf("]\n")

	f.Printf("Segundo vetor resultante: [ ")
	for _, v := range vetresult2 {
		f.Printf("%d ", v)
	}
	f.Printf("]\n")
}
