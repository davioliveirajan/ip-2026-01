package main

import f"fmt"

func main() {
	var base, expoente int
	f.Printf("Digite o número da base: ")
	f.Scan(&base)
	f.Printf("Digite o número do expoente: ")
	f.Scan(&expoente)

	f.Printf("O resultado é %d\n", potencia(base, expoente))

}

func potencia(base, expoente int) int { 
	if expoente == 0 {
		return 1
	} else {
		return base * potencia(base, expoente-1)
	}
}