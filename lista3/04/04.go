package main

import f"fmt"

func main() {
	var pot int

	f.Printf("Digite um número: ")
	f.Scan(&pot)

	i := 1

	for {

		if i * i == pot {
			f.Printf("A raiz quadrada de %d é %d\n", pot, i)
			f.Printf("Digite um número: ")
			f.Scan(&pot)
			i = 1
		} else {
			i++
		}

		if i > pot {
			f.Printf("%d não é um quadrado perfeito.\n", pot)
			f.Printf("Digite um número: ")
			f.Scan(&pot) 
			i =1

		if pot <= 0 {
		f.Println("A potência deve ser um número inteiro não negativo.")
		break
		}
		}
	}
	}