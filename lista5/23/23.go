package main

import f"fmt"

func main() {

	var janela [24]int
	var corredor [24]int

	for {
		f.Println("\n===== VENDA DE PASSAGENS =====")
		f.Println("1 - Poltrona na janela")
		f.Println("2 - Poltrona no corredor")
		f.Println("3 - Encerrar")
		f.Print("Escolha: ")

		var op int
		f.Scan(&op)

		if op == 3 {
			f.Println("Programa encerrado.")
			break
		}

		totalOcupadas := 0

		for i := 0; i < 24; i++ {
			totalOcupadas += janela[i]
			totalOcupadas += corredor[i]
		}

		if totalOcupadas == 48 {
			f.Println("Ônibus completamente cheio!")
			break
		}

		switch op {

		case 1:
			f.Println("\nPoltronas da janela disponíveis:")

			existe := false

			for i := 0; i < 24; i++ {
				if janela[i] == 0 {
					f.Printf("%d ", i)
					existe = true
				}
			}

			if !existe {
				f.Println("\nNão existem poltronas livres na janela.")
				continue
			}

			var poltrona int
			f.Print("\nEscolha a poltrona: ")
			f.Scan(&poltrona)

			if poltrona >= 0 && poltrona < 24 {
				if janela[poltrona] == 0 {
					janela[poltrona] = 1
					f.Println("Passagem vendida com sucesso!")
				} else {
					f.Println("Poltrona já ocupada.")
				}
			} else {
				f.Println("Poltrona inválida.")
			}

		case 2:
			f.Println("\nPoltronas do corredor disponíveis:")

			existe := false

			for i := 0; i < 24; i++ {
				if corredor[i] == 0 {
					f.Printf("%d ", i)
					existe = true
				}
			}

			if !existe {
				f.Println("\nNão existem poltronas livres no corredor.")
				continue
			}

			var poltrona int
			f.Print("\nEscolha a poltrona: ")
			f.Scan(&poltrona)

			if poltrona >= 0 && poltrona < 24 {
				if corredor[poltrona] == 0 {
					corredor[poltrona] = 1
					f.Println("Passagem vendida com sucesso!")
				} else {
					f.Println("Poltrona já ocupada.")
				}
			} else {
				f.Println("Poltrona inválida.")
			}

		default:
			f.Println("Opção inválida.")
		}
	}
}