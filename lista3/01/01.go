package main

import f "fmt"

func main() {
	var base, exp int
	f.Print("Digite a base: ")
	f.Scan(&base)
	f.Print("Digite o expoente: ")
	f.Scan(&exp)

	potencia := 1
	for exp > 0 {
		potencia *= base
		exp--
	}
	f.Printf("O resultado é: %d\n", potencia)
} 
