package main

import f"fmt"

func main() {
	var vetor[] int
	var cod, n int

	f.Scan(&cod)
	for i := 0; i < 10; i++ {
		f.Scan(&n)
		vetor = append(vetor, n)
	}

	switch cod{
	case 0: 
		return
	case 1:
		f.Println(vetor)
	case 2:
		f.Printf("[ ")
		for i := 9; i >= 0; i-- {
			f.Printf("%d ", vetor[i])
		}
		f.Printf("]")
	}
}
