package main

import "fmt"

func main() {
	var (d, m, a int
		mes string)
	fmt.Printf("Digite o dia: ")
	fmt.Scan(&d)
	fmt.Printf("Digite o mês: ")
	fmt.Scan(&m)
	fmt.Printf("Digite o ano: ")
	fmt.Scan(&a)

	if d > 31 || m > 12 || a < 0 {
		fmt.Printf("Data inválida\n")
	} else {
		fmt.Printf("Data válida\n")
	}

	switch m {
	case 1:
		mes = "Janeiro"
	case 2:
		mes = "Fevereiro"
	case 3:
		mes = "Março"
	case 4:
		mes = "Abril"
	case 5:
		mes = "Maio"
	case 6:
		mes = "Junho"
	case 7:
		mes = "Julho"
	case 8:
		mes = "Agosto"
	case 9:
		mes = "Setembro"
	case 10:
		mes = "Outubro"
	case 11:
		mes = "Novembro"
	case 12:
		mes = "Dezembro"
	}
	fmt.Printf("Data: %d de %s de %d\n", d, mes, a)
}