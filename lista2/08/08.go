package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Printf("Digite um número: ")
	fmt.Scan(&x)

	if x > 20 && x < 90 {
		fmt.Printf("O número está entre 20 e 90\n")
	} else {
		fmt.Printf("O número não está entre 20 e 90\n")
	}
}
