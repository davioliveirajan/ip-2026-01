package main

import "fmt"

func main() {
	var n1, n2, n3, q int
	fmt.Printf("Digite 3 números:")
	fmt.Scanf("%d%d%d", &n1, &n2, &n3)

	q = n1*100 + n2*10 + n3

	if n1 > 9 || n2 > 9 || n3 > 9 {
		fmt.Printf("DIGITO INVÁLIDO")
	} else {
		fmt.Printf("%d, %d", q, q*q)
	}

}
