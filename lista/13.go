package main

import "fmt"

func main() {
	var (
		n        float32
		conceito string
	)
	fmt.Printf("Nota: ")
	fmt.Scan(&n)

	if n >= 9 {
		conceito = "A"
	} else if n < 9 && n >= 7.5 {
		conceito = "B"
	} else if n < 7.5 && n >= 6 {
		conceito = "C"
	} else {
		conceito = "D"
	}

	fmt.Printf("NOTA = %.1f CONCEITO = %s\n", n, conceito)
}
