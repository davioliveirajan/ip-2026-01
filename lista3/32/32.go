package main

import f"fmt"

func main() {
	var x, y, mult float64

	f.Printf("Digite o valor dos números que quer multiplicar: ")
	f.Scan(&x, &y)

	for i := 1; float64(i) <= y; i++ {
		mult += x
	}

	f.Printf("O resultado da multiplicação é: %.2f\n", mult)
}
