package main

import f"fmt"

func main() {
	var ai, a1, a2, n int
	f.Printf("Digite a quantidade termos da sequência: ")
	f.Scan(&n)

	if n < 3 {
		f.Printf("O número deve ser maior que 3\n")
		return
	}

	f.Printf("Digite o primeiro termo e o segundo termo: ")
	f.Scan(&a1, &a2)

	f.Printf("%d\n", a1)
	f.Printf("%d\n", a2)

	for i := 3; i <= n; i++ {
		if i % 2 == 1 {
			ai = a2 + a1
			f.Printf("%d\n", ai)
			a1 = a2
			a2 = ai
		} else {
			ai = a2 - a1
			f.Printf("%d\n", ai)
			a1 = a2
			a2 = ai
		}
	}
}
