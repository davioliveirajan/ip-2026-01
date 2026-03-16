package main

import (
	"fmt"
)

func main() {

	var sm, cn, cs, cc, d float32
	fmt.Printf("Valor do salário mínimo: ")
	fmt.Scan(&sm)
	fmt.Printf("Quantidade de kW da residência: ")
	fmt.Scan(&cn)

	cs = sm * 0.7 / 100
	cc = cs * cn
	d = cc * 0.9

	fmt.Printf("Custo por kW: R$ %.2f\n", cs)
	fmt.Printf("Custo do consumo: R$ %.2f\n", cc)
	fmt.Printf("Custo com desconto: R$ %.2f", d)
}
