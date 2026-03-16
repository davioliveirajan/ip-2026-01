package main

import "fmt"

func main() {
	var n, s float64

	fmt.Printf("Digite um número inteiro e positivo: ")
	fmt.Scan(&n)

	for n > 0 {
		s = s + 1/n
		n--
	}
	fmt.Printf("%f\n", s)

}
