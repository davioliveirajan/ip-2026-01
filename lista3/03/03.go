package main

import f"fmt"

func main() {
	var carlos, joao, m float64
	f.Print("Digite o salário de Carlos: ")
	f.Scan(&carlos)

	joao = carlos / 3

	for {
		carlos = carlos * 1.02
		joao = joao * 1.05
		m++
		
		if joao >= carlos {
			break
		}
}
	f.Printf("Serão necessários %.0f meses para que João receba o mesmo salário que Carlos.\n", m)
}