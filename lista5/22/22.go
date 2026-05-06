package main

import "fmt"

func main() {

	var codigos [10]int
	var saldos [10]float64

	for i := 0; i < total; i++ {
		for {
			fmt.Printf("Digite o código da conta %d: ", i+1)
			fmt.Scan(&codigos[i])

			existe := false
			for j := 0; j < i; j++ {
				if codigos[j] == codigos[i] {
					existe = true
					break
				}
			}

			if existe {
				fmt.Println("Código já cadastrado! Digite outro.")
			} else {
				break
			}
		}

		fmt.Printf("Digite o saldo da conta %d: ", i+1)
		fmt.Scan(&saldos[i])
	}

	for {
		fmt.Println("\n1. Depósito")
		fmt.Println("2. Saque")
		fmt.Println("3. Ativo bancário")
		fmt.Println("4. Sair")
		fmt.Print("Escolha: ")

		var op int
		fmt.Scan(&op)

		switch op {
		case 1:
			var codigo int
			var valor float64

			fmt.Print("Código da conta: ")
			fmt.Scan(&codigo)

			pos := -1
			for i := 0; i < total; i++ {
				if codigos[i] == codigo {
					pos = i
					break
				}
			}

			if pos == -1 {
				fmt.Println("Conta não encontrada")
				continue
			}

			fmt.Print("Valor do depósito: ")
			fmt.Scan(&valor)
			saldos[pos] += valor

		case 2: 
			var codigo int
			var valor float64

			fmt.Print("Código da conta: ")
			fmt.Scan(&codigo)

			pos := -1
			for i := 0; i < total; i++ {
				if codigos[i] == codigo {
					pos = i
					break
				}
			}

			if pos == -1 {
				fmt.Println("Conta não encontrada")
				continue
			}

			fmt.Print("Valor do saque: ")
			fmt.Scan(&valor)

			if saldos[pos] >= valor {
				saldos[pos] -= valor
			} else {
				fmt.Println("Saldo insuficiente")
			}

		case 3: 
			var totalSaldo float64
			for i := 0; i < 10; i++ {
				totalSaldo += saldos[i]
			}
			fmt.Printf("Ativo bancário: %.2f\n", totalSaldo)

		case 4: 
			fmt.Println("Encerrando...")
			return

		default:
			fmt.Println("Opção inválida")
		}
	}
}
