package main

import f"fmt"

func main() {
	var n int
	f.Printf("Digite um número: ")
	f.Scan(&n)

	for i := 1; i <= n; i++ {
		f.Printf("%d\n", i*i)
	}
}
