package main

import f"fmt"

func main() {
	var n1, n2, q, r int

	f.Printf("Digite o numerador: ")
	f.Scan(&n1)
	f.Printf("Digite o denominador: ")
	f.Scan(&n2)

	for n1 > n2 {
		n1 -= n2
		q++
	}
	r = n1

	f.Printf("O quociente é: %d\n", q)
	f.Printf("O resto é: %d\n", r)
}