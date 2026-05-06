package main

import "fmt"

func main() {
	var nota int


	freqAbs := make([]int, 11)

	for i := 0; i < 15; i++ {
		fmt.Printf("Digite a %dª nota (0 a 10): ", i+1)
		fmt.Scan(&nota)

		if nota >= 0 && nota <= 10 {
			freqAbs[nota]++
		} else {
			fmt.Println("Nota inválida! Digite novamente.")
			i--
		}
	}

	fmt.Println("\nNota | Freq. Absoluta | Freq. Relativa")
	fmt.Println("--------------------------------------")

	for i := 0; i <= 10; i++ {
		freqRel := float64(freqAbs[i]) / float64(15)
		fmt.Printf("%4d | %15d | %14.2f\n", i, freqAbs[i], freqRel)
	}
}