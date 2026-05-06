package main 

import f"fmt"

func main() {
	var n[] int
	var menor, n1, indice int
	f.Printf("Digite 10 números: ")
	for i:= 0; i < 10; i++ {
		f.Scan(&n1)
		n = append(n, n1)
	}

	for i, v := range n {
		if i == 0 {
			menor = v
			indice = 1
		} else if v < menor {
			menor = v
			indice = i + 1
		}
	}

	f.Printf("O menor número do vetor é %d e sua posição é %d\n", menor, indice)
}
