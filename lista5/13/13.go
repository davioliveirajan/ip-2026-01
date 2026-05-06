package main

import f"fmt"

func main() {
	var meses[] int
	var n, m, recen1, recen2, recen3, ind1, ind2, ind3 int

	i := 1

	for {
		f.Printf("%d empregado: ", i)
		f.Scan(&n, &m)

		meses = append(meses, m)

		if n == 0 && m == 0 {
			break
		}
		i++
	}

	for i, v := range meses {
		if recen1 == 0 {
			recen1 = v
			ind1 = i
		} else if v > recen1 {
			recen3 = recen2
			recen2 = recen1
			recen1 = v
			ind3 = ind2
			ind2 = ind3
			ind1 = i
		}

	}

	f.Printf("Mais recente: N°%d, %d meses", ind1, recen1)
	f.Printf("2º mais recente: N°%d, %d meses", ind2, recen2)
	f.Printf("3º mais recente: N°%d, %d meses", ind3, recen3)

}
