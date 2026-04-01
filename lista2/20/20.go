package main

import "fmt"

func main() {
	var preco, pag int
	fmt.Printf("Digite o preço do produto: ")
	fmt.Scan(&preco)
	fmt.Printf("Digite a condição de pagamento: (1-4)")
	fmt.Scan(&pag)
	switch pag {
	case 1:
		preco = preco * 0.9
	case 2:
		preco = preco * 0.95
	case 4:
		preco = preco * 1.1
	}

	fmt.Printf("O preço final é %.2f\n", preco)
}