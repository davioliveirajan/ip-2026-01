package main

import f "fmt"

func main() {
	var n, a, e, r int
	var n1, n2, soma, media, media_sala float64

	f.Printf("Digite a quantidade de alunos: ")
	f.Scan(&n)

	for i := 0; i < n; i++ {
		f.Printf("Digite a nota 1: ")
		f.Scan(&n1)

		if n1 < 0 || n1 > 10 {
			f.Printf("Nota inválida. Digite uma nota entre 0 e 10.\n")
			i--
			continue
		}

		f.Printf("Digite a nota 2: ")
		f.Scan(&n2)

		if n2 < 0 || n2 > 10 {
			f.Printf("Nota inválida. Digite uma nota entre 0 e 10.\n")
			i--
			continue
		}

		soma = n1 + n2
		media = soma / 2

		f.Printf("Média do aluno %d: %.2f\n", i+1, media)

		if media >= 7 {
			f.Printf("Aprovado\n")
			a++
		} else if media >= 3 {
			f.Printf("Exame\n")
			e++
		} else {
			f.Printf("Reprovado\n")
			r++
		}
		
		media_sala += media
	}

	media_sala /= float64(n)
	f.Printf("Quantidade de alunos aprovados: %d\n", a)
	f.Printf("Quantidade de alunos em exame: %d\n", e)
	f.Printf("Quantidade de alunos reprovados: %d\n", r)
	f.Printf("Média da sala: %.2f\n", media_sala)
}