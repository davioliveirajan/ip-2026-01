package main 

import "fmt"

func main() {
	var d, preco, cat float64
	fmt.Printf("Digite o preço normal do DVD: ")
	fmt.Scan(&preco)
	fmt.Printf("Digite o dia da semana (1-7): ")
	fmt.Scan(&d)
	fmt.Printf("A categoria do DVD é: (n/l)")
	fmt.Scan(&cat)

	switch cat {
		case 'l':
			preco = preco * 1.15
	}

	switch d {
	case 2, 4, 5:
		preco = preco * 0.6
	}

	fmt.Printf("O preço final do DVD é: %.2f\n", preco)

}
