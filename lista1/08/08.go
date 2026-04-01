package main

import "fmt"

func main() {
	var r, h, custo, cir, lat float32
	fmt.Printf("Raio da lata: ")
	fmt.Scan(&r)
	fmt.Printf("Altura da lata: ")
	fmt.Scan(&h)

	cir = 3.14159 * r * r
	lat = 2 * 3.14159 * r * h
	custo = 100 * (2*cir + lat)

	fmt.Printf("O valor do custo é = %.2f\n", custo)
}
