package main

import "fmt"
import "math"

func main() {
	var (fig int
		h, r, v, a float64)
	fmt.Printf("Digite a figura desejada (1-3): ")
	fmt.Scan(&fig)

	switch fig {
	case 1: 
	v = 3.14 * r * r * h / 3
	a = 3.14 * r * math.Sqrt(r*r+h*h)
	case 2:
	v = 3.14 * r * r * h
	a = 2 * 3.14 * r * h
	case 3:
	v = 4 * 3.14 * r * r * r / 3
	a = 4 * 3.14 * r * r
	}
	fmt.Printf("Digite o raio e a altura: ")
	fmt.Scan(&r, &h)
	fmt.Printf("O volume é %.2f e a área é %.2f\n", v, a)
}