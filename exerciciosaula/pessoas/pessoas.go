package main

import f "fmt" 

type Pessoa struct {
	nome string
	altura float64
	pesoIdeal float64
}
func main() {
	pessoas := make([]Pessoa, 10)
	var total int

	for i := 0; i < len(pessoas); i++ {
		f.Printf("Digite o nome da pessoa %d: ", i + 1)
		f.Scan(&pessoas[i].nome)

		if pessoas[i].nome == "FIM" {
			break
		}
		f.Printf("Digite a altura da pessoa %d: ", i + 1)
		f.Scan(&pessoas[i].altura)
		total++
	}
	for n := 0; n < total; n++ {
		pessoas[n].pesoIdeal = (pessoas[n].altura * 72.7) - 58
		f.Printf("Pessoa %d:\n Nome: %s\n Altura: %.2f\n Peso Ideal: %.2f\n", n + 1,  pessoas[n].nome, pessoas[n].altura, pessoas[n].pesoIdeal)
	}
}