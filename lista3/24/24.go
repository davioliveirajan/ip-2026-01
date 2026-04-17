package main

import f"fmt"

func main() {
	var sen, a float64
	var A[]float64
	var angulos[]float64

	a = 0

	for i := a; i <= 6.3; i += 0.1 {
		angulos = append(angulos, a)

		sen = a - (a * a * a / 6) + (a * a * a * a * a / 120) - (a * a * a * a * a * a * a / 5040)
		A = append(A, sen)
		a += 0.1
	}


		for n, i := range A {
			f.Printf("Seno de %.2f = %.2f\n", angulos[n], i)
		}
	}
