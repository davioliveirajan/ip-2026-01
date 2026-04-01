package main

import "fmt"

func main() {
	var n1, n2, n3, ex, media int
	fmt.Printf("Digite a primeira nota: ")
	fmt.Scan(&n1)
	fmt.Printf("Digite a segunda nota: ")
	fmt.Scan(&n2)
	fmt.Printf("Digite a terceira nota: ")
	fmt.Scan(&n3)
	fmt.Printf("Digite a nota média dos exercícios: ")
	fmt.Scan(&ex)

	media = n1 + n2 * 2 + n3 * 3 + ex / 7

	fmt.Printf("A média final é %d\n", media)
}