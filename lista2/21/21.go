package main

import "fmt"

func main() {
	var n1, n2, n3, ex, media float64
	var conc string
	fmt.Printf("Digite a primeira nota: ")
	fmt.Scan(&n1)
	fmt.Printf("Digite a segunda nota: ")
	fmt.Scan(&n2)
	fmt.Printf("Digite a terceira nota: ")
	fmt.Scan(&n3)
	fmt.Printf("Digite a nota média dos exercícios: ")
	fmt.Scan(&ex)

	media = n1 + n2 * 2 + n3 * 3 + ex / 7

	fmt.Printf("A média final é %.2f\n", media)

	if media >= 9 {
		conc = "A"
	} else if media >= 7.5 {
		conc = "B"
	} else if media >= 6 {
		conc = "C"
	} else if media >= 4 {
		conc = "D"
	} else {
		conc = "E"
	}

	if media >= 6 {
		fmt.Printf("O conceito é %s, aprovado\n", conc)
	} else {
		fmt.Printf("O conceito é %s, reprovado\n", conc)
	}
}