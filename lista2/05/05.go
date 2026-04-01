package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Digite um número: ")
	fmt.Scan(&n)

	if n%5 == 0 {
		fmt.Printf("Divisível por 5\n")
	} else {
		fmt.Printf("Não divisível por 5\n")
	}
}
