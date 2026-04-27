package main

import f"fmt"

func main() {
	var n int
	f.Printf("Digite um número inteiro: ")
	f.Scanf("%d", &n)
	resultado := binario(n)
	f.Printf("Número em binário: %d\n", resultado)
}

func binario(n int) int {
	if n == 0 {
		return 0
	}
	return n%2 + 10*binario(n/2)
}
