package main 

import f"fmt"

func main() {
	var vet[] int
	var pares[] int
	var impares[] int
	var n, p, i int
	f.Printf("Digite 10 números: ")
	for i := 0; i < 10; i++ {
		f.Scan(&n)
		vet = append(vet, n)
	}

	for _, v := range vet {
		if v % 2 == 0 {
			pares = append(pares, v)
			p++
		} else {
			impares = append(impares, v)
			i++
		}
	}
	f.Printf("Pares digitados: ")
	for _, v := range pares {
		f.Printf("%d ", v)
	}
	f.Println()
	f.Printf("Quantidade de pares: %d\n", p)

	f.Printf("Ímpares digitados: ")
	for _, v := range impares {
		f.Printf("%d ", v)
	}
	f.Println()
	f.Printf("Quantidade de ímpares: %d\n", i)
}