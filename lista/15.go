package main

import "fmt"

func main() {
	var n, q int
	fmt.Printf("Digite um número entre 5 e 2000: ")
	fmt.Scan(&n)

	if n%2 > 0 {
		for n-1 > 0 {
			q = (n - 1) * (n - 1)
			fmt.Printf("%d^2 = %d\n", (n - 1), q)
			n -= 2
		}
	} else {
		for n > 0 {
			q = n * n
			fmt.Printf("%d^2 = %d\n", n, q)
			n -= 2
		}
	}
}
