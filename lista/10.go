package main

import "fmt"

func main() {
	var a, b, c, d, det float32
	fmt.Printf("Elemento a: ")
	fmt.Scan(&a)
	fmt.Printf("Elemento b: ")
	fmt.Scan(&b)
	fmt.Printf("Elemento c: ")
	fmt.Scan(&c)
	fmt.Printf("Elemento d: ")
	fmt.Scan(&d)

	det = a*d - b*c

	fmt.Printf("O valor do detrminante é: %.2f\n", det)
}
