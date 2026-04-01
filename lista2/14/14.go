package main

import "fmt"

func main() {
	var (
		preco int
		ar, pm, ve, dh string
	)
	fmt.Printf("Digite o preço do carro: ")
	fmt.Scan(&preco)
	fmt.Printf("Adicionar ar condicionado? (s/n): ")
	fmt.Scan(&ar)
	fmt.Printf("Adicionar pintura metálica? (s/n): ")
	fmt.Scan(&pm)
	fmt.Printf("Adicionar vidro elétrico? (s/n): ")
	fmt.Scan(&ve)
	fmt.Printf("Adicionar direção hidráulica? (s/n): ")
	fmt.Scan(&dh)

	if ar == "s" {
		preco = preco + 1750
	}
	if pm == "s" {
		preco = preco + 800
	}
	if ve == "s" {
		preco = preco + 1200
	}
	if dh == "s" {
		preco = preco + 2000
	}
	fmt.Printf("Preço final do carro: %d\n", preco)
}
