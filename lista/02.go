package main

import "fmt"

func main() {
	var (
		pub, pop, ger, arq, cad, renda float32
		jogos                          int
	)
	fmt.Printf("Qual o número de jogos: ")
	fmt.Scan(&jogos)

	for n := 1; n <= jogos; n++ {
		fmt.Printf("Publico e porcentagens: ")
		fmt.Scan(&pub, &pop, &ger, &arq, &cad)

		renda = (pop + ger*5 + arq*10 + cad*20) * pub / 100

		fmt.Printf("A renda do jogo %d é %.2f\n", n, renda)
	}
}
