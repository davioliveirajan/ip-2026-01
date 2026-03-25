package main

import (
	"fmt"
	"math"
)

func main() {
	var n float32
	fmt.Printf("Digite um número: ")
	fmt.Scan(&n)

	rq := math.Sqrt(float64(n))

	fmt.Printf("Raiz Quadrada de %.f = %f\n", n, rq)
}
