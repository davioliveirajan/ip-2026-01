package main

import f "fmt"

func main() {
	var s, x float64

	f.Printf("Digite um número: ")
	f.Scan(&x)

	for i := 0; i <= 10; i += 2 {

		s += x / fatorial(float64(i)) - x / fatorial(float64(i+1))

}
	f.Printf("Resultado: %.6f\n", s)
}

func fatorial(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}