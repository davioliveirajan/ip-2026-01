package main 

import f"fmt"

func main() {
	var p, magro, gordo float64
	var ng, nm int
	for i := 1; i <= 10; i++ {
		f.Printf("Digite o peso do boi número %d: ", i)
		f.Scan(&p)

		if p > gordo || magro == 0 {
			gordo = p
			ng = i
		}

		if p < magro || magro == 0 {
			magro = p
			nm = i

		}
	}

	f.Printf("Boi mais gordo =\n")
	f.Printf("Número %d\n", ng)
	f.Printf("Peso = %.2f kg\n\n", gordo)
	f.Printf("Boi mais magro =\n")
	f.Printf("Número %d\n", nm)
	f.Printf("Peso = %.2f kg\n", magro)
}
