package main

import f"fmt"

func main() {
	var idades[] int
	var age, q, max, num int

	for i := 0; i < 10; i++ {
		f.Scan(&age)
		idades = append(idades, age)
	}

	for i, v := range idades {
		q = 1
		for j, l := range idades {
			if i == j {

			} else if v == l {
				q++
			}
			if q == 1 || q > max {
				max = q
				num = l
			package main

import "fmt"

func main() {
	const total = 15
	var nota int

	// vetor para frequência absoluta (índices de 0 a 10)
	freqAbs := make([]int, 11)

	// leitura das notas
	for i := 0; i < total; i++ {
		fmt.Printf("Digite a %dª nota (0 a 10): ", i+1)
		fmt.Scan(&nota)

		if nota >= 0 && nota <= 10 {
			freqAbs[nota]++
		} else {
			fmt.Println("Nota inválida! Digite novamente.")
			i-- // repete a leitura
		}
	}

	// impressão da tabela
	fmt.Println("\nNota | Freq. Absoluta | Freq. Relativa")
	fmt.Println("--------------------------------------")

	for i := 0; i <= 10; i++ {
		freqRel := float64(freqAbs[i]) / float64(total)
		fmt.Printf("%4d | %15d | %14.2f\n", i, freqAbs[i], freqRel)
	}
}package main

import "fmt"

func main() {
	const total = 15
	var nota int

	// vetor para frequência absoluta (índices de 0 a 10)
	freqAbs := make([]int, 11)

	// leitura das notas
	for i := 0; i < total; i++ {
		fmt.Printf("Digite a %dª nota (0 a 10): ", i+1)
		fmt.Scan(&nota)

		if nota >= 0 && nota <= 10 {
			freqAbs[nota]++
		} else {
			fmt.Println("Nota inválida! Digite novamente.")
			i-- // repete a leitura
		}
	}

	// impressão da tabela
	fmt.Println("\nNota | Freq. Absoluta | Freq. Relativa")
	fmt.Println("--------------------------------------")

	for i := 0; i <= 10; i++ {
		freqRel := float64(freqAbs[i]) / float64(total)
		fmt.Printf("%4d | %15d | %14.2f\n", i, freqAbs[i], freqRel)
	}
}
		}
	}

	f.Printf("Moda = %d, aparece %d vezes\n", num, max)
}