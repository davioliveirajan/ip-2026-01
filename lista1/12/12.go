package main

import "fmt"

func main() {
	var h, valor int

	fmt.Printf("Número de horas utilizadas: ")
	fmt.Scan(&h)

	if h%3 > 0 {
		valor = (h-h%3)*10/3 + h%3*5
	}

	fmt.Printf("O valor a pagar é: %.2d", valor)
}
