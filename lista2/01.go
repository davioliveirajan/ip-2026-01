
package main

import "fmt"

func main() {
	var n int
	fmt.Printf("Digite um número: ")
	fmt.Scan(&n)

	if n%2 == 1 {
		fmt.Printf("Ímpar\n")
	} else {
		fmt.Printf("Par\n")
	}
}
