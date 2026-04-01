package main

import "fmt"

func main() {
	var n1, n2, soma int
	fmt.Printf("Digite dois números inteiros: ")
	fmt.Scan(&n1, &n2)

	soma = n1 + n2

	if soma > 20 {
		soma += 8
	} else {
		soma -= 5
	}
	fmt.Printf("Soma = %d\n", soma)
}
