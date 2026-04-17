package main

import f"fmt"

func main() {
	var cpf string
	f.Print("Digite o CPF: ")
	f.Scan(&cpf)

	if int(cpf[9] - '0') != pdigito(cpf) {
		f.Println("CPF inválido")
	} else {
		if int(cpf[10] - '0') != sdigito(cpf) {
			f.Println("CPF inválido")
		} else {
			f.Println("CPF válido")
		}
		
	}

}

func pdigito(cpf string) int {
	var d int

	for i := 0; i < 9; i++ {
		d += int(cpf[i] - '0') * (10 - i)
	}

	if d % 11 < 2 {
		d = 0
		return d
	} else {
		d = 11 - (d % 11)
		return d
	}
}

func sdigito(cpf string) int {
	var d int
	
	for i := 0; i < 10; i++ {
		d += int(cpf[i] - '0') * (11 - i)
	}

	d += pdigito(cpf) * 2

	if d % 11 < 2 {
		d = 0
		return d
	} else {
		d = 11 - (d % 11)
		return d
	}
}