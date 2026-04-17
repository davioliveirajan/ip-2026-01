package main

import f"fmt"

func main() {
	var soma, n int
	f.Printf("Digite um número inteiro positivo: ")
	f.Scan(&n)

	for i := 1; i <= n; i++ {
		soma += i
	}

	f.Printf("A soma dos números de 1 a %d é: %d\n", n, soma)
}
