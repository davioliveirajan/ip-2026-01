package main

import "fmt"
import "math"

func main() {
	var a, b, c, delta, x1, x2 float64
	fmt.Printf("Digite o valor de a: ")
	fmt.Scan(&a)
	fmt.Printf("Digite o valor de b: ")
	fmt.Scan(&b)
	fmt.Printf("Digite o valor de c: ")
	fmt.Scan(&c)
	delta = b*b - 4*a*c
	if delta < 0 {
		fmt.Printf("Raízes imaginárias\n")
	} else if delta == 0 {
		x1 = -b / (2 * a)
		fmt.Printf("Raiz única: %.2f\n", x1)
	} else {
		x1 = (-b + math.Sqrt(delta)) / (2 * a)
		x2 = (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("Raízes distintas: %.2f e %.2f\n", x1, x2)
	}
}
