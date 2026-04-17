package main

import f "fmt"

func main() {
	var n, i, idade, altura, somaa, peso, a, p, x, total, percentual, media float64



	for {

		f.Printf("Digite a idade: ")
		f.Scan(&idade)

		if idade > 50 {
			i++
		}

		f.Printf("Digite a altura: ")
		f.Scan(&altura)

		if idade >= 10 && idade <= 20 {
			somaa += altura
			a++
		}

		media = somaa / a

		f.Printf("Digite o peso: ")
		f.Scan(&peso)

		if peso < 40 {
			p++
		} else {
			x++
		}

		f.Printf("Deseja continuar? 1- Sim, 0- Não: ")
		f.Scan(&n)

		if n == 0 {
			f.Printf("Número de pessoas com mais de 50 anos: %.0f\n", i)
			f.Printf("Média das alturas das pessoas entre 10 e 20 anos: %.2f\n", media)
			f.Printf("Percentual de pessoas com peso inferior a 40kg: %.2f%%\n", percentual)
			break
		} else if n != 1 {
			f.Println("Opção inválida. Digite 1 para Sim ou 0 para Não.")
			continue
		}

		total = p + x
		percentual = p / (total + 1) * 100
	}
}
