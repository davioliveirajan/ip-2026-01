package main

import f "fmt"

func main() {
	var n, soma, pares, p, imp, maior, menor int
	var media, mediap, perc float64

	i := 0
	maior = n
	menor = 9223372036854775807

	for {
		f.Printf("Digite um número: ")
		f.Scan(&n)
		if n == 30000 {
			break
		}

		if  n > maior {
			maior = n
		}
		
		if n < menor {
			menor = n
		}

		if n % 2 == 0 {
			pares += n
			p++

		} else {
			imp++
		}

		soma = soma + n
		i++

	}

	media = float64(soma) / float64(i)
	perc = (float64(imp) * 100) / float64(i)
	mediap = float64(pares) / float64(p)

	f.Printf("Soma: %d\n", soma)
	f.Printf("Média: %.2f\n", media)
	f.Printf("Maior número: %d\n", maior)
	f.Printf("Menor número: %d\n", menor)
	f.Printf("Média dos pares: %.2f\n", mediap)
	f.Printf("Percentual de ímpares: %.2f%%\n", perc)
}