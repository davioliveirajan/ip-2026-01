package main

import "fmt"

func main() {
	var idade int
	fmt.Printf("Digite sua idade: ")
	fmt.Scan(&idade)

	if idade < 16 {
		fmt.Println("Não-eleitor")
	} else if idade < 18 || idade >= 65 {
		fmt.Println("Eleitor facultativo")
	} else {
		fmt.Println("Eleitor obrigatório")
	}
}
		
