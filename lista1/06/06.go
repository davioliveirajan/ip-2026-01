package main

import "fmt"

func main() {
	var (
		n    int
		f, c float32
	)
	fmt.Printf("Número de entradas: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Temperature em Fahrenheit: ")
		fmt.Scan(&f)

		c = (f - 32) * 5 / 9

		fmt.Printf("%.2f Fahrenheit Equivale a %.2f Celsius\n", f, c)
	}

}
