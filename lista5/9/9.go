package main 

import f"fmt"

func main() {
	var vetor[] float64
	var media, altura, soma float64

	f.Printf("Digite a altura de 10 jogadores: ")
	for i := 0; i< 10; i++ {
		f.Scan(&altura)
		vetor = append(vetor, altura)

		soma += altura
	}

	media = soma / 10

	for j, i := range vetor {
		if i > media {
			f.Printf("%.2f m ", vetor[j])
		}
	}
}
