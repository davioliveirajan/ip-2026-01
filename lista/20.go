package main

import "fmt"

func main() {
	var h, m, s, t int
	fmt.Printf("Valor em horas: ")
	fmt.Scan(&h)
	fmt.Printf("Valor em minutos: ")
	fmt.Scan(&m)
	fmt.Printf("Valor em segundos: ")
	fmt.Scan(&s)

	t = h*3600 + m*60 + s

	fmt.Printf("O tempo em segundo é = %d\n", t)
}
