package main

import "fmt"

func main() {
	var a, b, c, l, s float32
	fmt.Printf("Digite 3 números: ")
	fmt.Scan(&a, &b, &c)

	l = a + b*(c-1)
	s = (a + l) * c / 2

	fmt.Printf("%.f\n", s)
}
