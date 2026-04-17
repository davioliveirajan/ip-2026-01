package main 

import f "fmt"

func main() {
	var n int
	f.Printf("Digite um número: ")
	f.Scan(&n)

	i := 1

	for {
		if i * (i + 1) * (i + 2) == n {
			f.Printf("%d é um número triangular\n", n)
			break
		} else {
			i++
		}
		if i > n {
			f.Printf("%d não é um número triangular\n", n)
			break
		}
	}
}
