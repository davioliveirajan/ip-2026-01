package main

import (
	"fmt"
)

func main() {
	var compra, venda float32
	fmt.Printf("Valor da compra: ")
	fmt.Scan(&compra)

	if compra < 10 {
		venda = compra * 1.7
	} else if compra >= 10 && compra < 30 {
		venda = compra * 1.5
	} else if compra >= 30 && compra < 50 {
		venda = compra * 1.4
	} else {
		venda = compra * 1.3
	}

	fmt.Printf("Valor da venda = %.2f\n", venda)
}
