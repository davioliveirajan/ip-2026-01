package main

import "fmt"

func main() {
	var n int
	fmt.Printf("Digite um número: ")
	fmt.Scan(&n)

	if n%3 > 0 && n%5 > 0 {
		fmt.Printf("O número não é divisível\n")
	} else {
		fmt.Printf("O número é divisível\n")
	}
}
