package main

import (
	"fmt"
)

func main() {
	var x float32
	fmt.Printf("Digite um número: ")
	fmt.Scan(&x)

	if x > 99 && x < 1000 {
		fmt.Printf("O número possui 3 dígitos\n")
	} else {
		fmt.Printf("O número não possui 3 dígitos\n")
	}
}
