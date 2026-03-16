package main

import (
	"fmt"
)

func main() {
	var (
		tipo                  string
		conta, consumo, preço float32
	)
	fmt.Printf("Conta, consumo e tipo: ")
	fmt.Scan(&conta, &consumo, &tipo)

	if tipo == "R" {
		preço = 5 + consumo/20
	} else if tipo == "C" {
		preço = 500 + (consumo-80)/4
	} else if tipo == "I" {
		preço = 800 + (consumo-100)/25
	} else {
		fmt.Printf("TIPO INVÁLIDO\n")
	}

	fmt.Printf("CONTA = %.f\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", preço)

}
