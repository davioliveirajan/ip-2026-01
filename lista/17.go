package main

import "fmt"

func main() {
	var x, y int
	fmt.Printf("Escreva dois números: ")
	fmt.Scan(&x, &y)

	if x%2 > 0 {
		fmt.Printf("O primeiro número não é par\n")
	} else {
		for y > 0 {
			fmt.Printf("%d\n", x)
			y--
			x += 2
		}
	}
}
