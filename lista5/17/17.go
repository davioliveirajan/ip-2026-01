package main 

import f"fmt"

func main() {
	var vetor[] int
	var primos[] int
	var n int 
	for i := 0; i < 10; i++ {
		f.Scan(&n)
		vetor = append(vetor, i)
	}

	for i, v := range vetor {
		if primo(vetor[i]) {
			primos = append(primos, v)
		}
	}

	f.Println(primos)

}

func primo(n int) bool {
	primo := 1

		if n < 2 {
			primo = 0
		} else {
			i := 2
			for i < n {
				if n%i == 0 {
					primo = 0
					break
				}
				i++
			}
		}

		if primo == 1 {
			return true
		} else {
			return false
		}
}
