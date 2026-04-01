package main

import (
	"fmt"
)

func main() {
	var x, y float32
	fmt.Printf("Digite o valor de x: ")
	fmt.Scan(&x)

	y = 8 / (2 - x)

	fmt.Printf("f(x) = %.1f\n", y)
}
