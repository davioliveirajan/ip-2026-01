package main

import f"fmt"
import "math"

func main() {
	var cos1, cos2, dif, x float64

	f.Printf("Digite um valor para x: ")
	f.Scanf("%f", &x)

	for i := 0; i < 20; i += 4 {
		cos1 += math.Pow(x, float64(i)) / float64(fatorial(i)) - math.Pow(x, float64(i+2)) / float64(fatorial(i+2))
	}

	cos2 = math.Cos(x)
	dif = cos1 - cos2

	f.Printf("%.2f\n", cos1)
	f.Printf("%.2f\n", dif)
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}