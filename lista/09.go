package main

import "fmt"

func main() {
	var a, b, c, delta float32
	fmt.Printf("Valor de A: ")
	fmt.Scan(&a)
	fmt.Printf("Valor de B: ")
	fmt.Scan(&b)
	fmt.Printf("Valor de C: ")
	fmt.Scan(&c)

	delta = b*b - 4*a*c

	fmt.Printf("O valor de delta é = %.2f\n", delta)
}
