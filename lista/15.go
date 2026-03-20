package main

import "fmt"

func main() {
	var n, q int
	fmt.Printf("Digite um número entre 5 e 2000: ")
	fmt.Scan(&n)

	if n%2 > 0 {
		for i := 1; i <= n; i += 2 {
			q = (i + 1) * (i + 1)
			fmt.Printf("%d^2 = %d\n", i+1, q)
		}
	} else {
		for i := 2; i <= n; i += 2 {
			q = i * i
			fmt.Printf("%d^2 = %d\n", i, q)
		}
	}
}
