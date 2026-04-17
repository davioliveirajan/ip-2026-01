package main

import f"fmt"

func main() {
	var preco, lucro, lucromax, i, a float64

	a = 0
	preco = 6

	for preco >= 1.2 {
		lucro = (preco * (130 + 30 * a)) - 300
		f.Printf("Preço: %.2f, Lucro: %.2f\n", preco, lucro)
		a++
		preco -= 0.6

		if lucro > lucromax || lucromax == 0 {
			lucromax = lucro
			i = (130 + 30 * a)
		}
	}
	f.Printf("Lucro máximo: %.2f\n", lucromax)
	f.Printf("Quantidade de ingressos: %.2f\n", i)
}