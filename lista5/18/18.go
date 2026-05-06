package main 

import f"fmt"

func main() {
	vetor := make([]int, 10)
	var n int

	for i := 0; i < 10; i++ {
		f.Scan(&n)
		j := i - 1
		for j >= 0 && vetor[j] > n {
			vetor[j + 1] = vetor[j]
			j--
		}
		vetor[j+1] = n
	}
	f.Println(vetor)
}
