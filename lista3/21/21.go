package main

import f"fmt"

func main() {
	var b, n, pot int
	f.Printf("Digite o valor da base: ")
	f.Scan(&b)

	if b < 2 {
		f.Printf("O valor da base deve ser maior ou igual a 2\n")
		return
	}

	f.Printf("Digite o valor do expoente: ")
	f.Scan(&n)

	if n <= 1 {
		f.Printf("O valor dop expoente deve ser maior que 1\n")
	}

	pot = 1

	for n > 0 {
		pot *= b
		n--
	}

	f.Printf("%d\n", pot)
}