package main

import f"fmt"

func main() {
	var a[] int
	var rep[] int
	var  quant[] int
	var n, q int

	f.Printf("Digite 10 números: ")
	for i := 0; i < 10; i++ {
		f.Scan(&n)
		a = append(a, n)
	}


	for i, v := range a {
		jaExiste := false
		for _, x := range rep {
			if x == v {
				jaExiste = true
				break
			}
		}

	if jaExiste {
		continue
	}

		q = 0
		for l, j := range a {
			if i == l {
				continue
			} else if v == j {
				q++
			}
		}

	if q > 0 {
		rep = append(rep, v)
		quant = append(quant, q+1)
	}

	j := len(rep)

	for i := 0; i < j; i++ {
		f.Printf("O número %d se repete %d vezes\n", rep[i], quant[i])
	}
}
}
