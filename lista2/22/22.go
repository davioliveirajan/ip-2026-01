package main

import "fmt"

func main() {
	var mat, sb, sl, inss, ir, hr float64
	fmt.Printf("Digite a matrícula: ")
	fmt.Scan(&mat)
	fmt.Printf("Digite o número de horas extra trabalhadas: ")
	fmt.Scan(&hr)

	sb = 3 * 788 + hr * 10

	if sb > 1500 {
		inss = sb * 0.12
	}

	if sb > 2000 {
		ir = sb * 0.2
	}

	sl = sb - inss - ir

	fmt.Printf("O salário líquido é %.2f\n", sl)
}

	
