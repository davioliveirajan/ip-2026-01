package main

import "fmt"

func main() {
	var s, sn float32
	fmt.Printf("Valor do salário: ")
	fmt.Scan(&s)

	if s <= 300 {
		sn = s * 1.5
	} else {
		sn = s * 1.3
	}
	fmt.Printf("Salário com reajuste = %.2f", sn)
}
