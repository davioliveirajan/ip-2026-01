package main

import f"fmt"

func main() {
	var x, y int 
	f.Printf("Digite o valor de x: ")
	f.Scan(&x)
	f.Printf("Digite o valor de y: ")
	f.Scan(&y)

	i := max(x, y)
	for {
		if x % i == 0 && y % i == 0 {
			f.Printf("O MDC de %d e %d é: %d\n", x, y, i)
			break
		} else {
			i--
		}
	}
}
