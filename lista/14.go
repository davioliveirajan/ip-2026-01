package main

import "fmt"

func main() {
	var h, a, b, v float32
	fmt.Printf("Altura e arestas do hexágono: ")
	fmt.Scan(&h, &a)

	b = 3 * a * a * 1.73205 / 2
	v = b * h / 3

	fmt.Printf("O volume da pirâmide é %.2f metros cúbicos", v)
}
