package main

import "fmt"

func main() {
	var a, b, c, media float32
	fmt.Printf("Nota das 3 provas:")
	fmt.Scanf("%f%f%f", &a, &b, &c)

	media = (a + b + c) / 3

	fmt.Printf("MÉDIA = %.2f\n", media)

	if media >= 6 {
		fmt.Printf("APROVADO\n")
	} else {
		fmt.Printf("REPROVADO\n")
	}

}
