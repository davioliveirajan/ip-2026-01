package main

import f "fmt"

func main() {
	var fibonacci, fibonacci1, fibonacci2, n int

	f.Printf("Digite um número maior que 3: ")
	f.Scan(&n)

	if n <= 3 {
		f.Printf("Número inválido. Digite um número maior que 3.\n")
		return
	}

	fibonacci1 = 0
	fibonacci2 = 1

	for i := 0; i <= n; i++ {
		f.Printf("%d\n", fibonacci1)
		fibonacci = fibonacci1 + fibonacci2
		fibonacci2 = fibonacci1
		fibonacci1 = fibonacci
	}
}
