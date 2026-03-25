package main

import "fmt"

func main() {
	var n int
	fmt.Printf("Digite um número: ")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Printf("Positivo\n")
	} else if n == 0 {
		fmt.Printf("Nulo\n")
	} else {
		fmt.Printf("Negativo\n")
	}
}
