package main

import "fmt"

func main() {
	var f, in, c, mm float32
	fmt.Printf("Digite a temperatura em F: ")
	fmt.Scan(&f)
	fmt.Printf("Digite a quantidade de chuva em polegadas: ")
	fmt.Scan(&in)

	c = (5*f - 160) / 9
	mm = in * 25.4

	fmt.Printf("O valor em Celsius: %.2f\n", c)
	fmt.Printf("A quantidade de chuva: %.2f\n", mm)
}
