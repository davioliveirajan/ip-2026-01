package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Printf("Digite 3 números: ")
	fmt.Scan(&n1, &n2, &n3)
	fmt.Printf("O maior número é %d\n", maior(n1, n2, n3))
}

func maior(n1, n2, n3 int) int {
	if n1 >= n2 && n1 >= n3 {
		return n1
	} else if n2 >= n3 {
		return n2
	} else {
		return n3
	}
}
