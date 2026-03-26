package main

import (
	"fmt"
)

func main() {
	var regiao, ida, preco float32
	fmt.Printf("Qual a região?")
	fmt.Scan(&regiao)
	fmt.Printf("Ida e volta?")
	fmt.Scan(&ida)

	if ida == 1 {
		switch regiao {
		case 1:
			preco = 900
		case 2:
			preco = 650
		case 3:
			preco = 600
		case 4:
			preco = 550
		}
	} else if ida == 2 {
		switch regiao {
		case 1:
			preco = 500
		case 2:
			preco = 350
		case 3:
			preco = 350
		case 4:
			preco = 300
		}
	}
	fmt.Printf("Preço = %.2f\n", preco)
}
