package main

import "fmt"

func main() {
	var l1, l2, l3 float32
	fmt.Printf("Medidas dos lados do triangulo: ")
	fmt.Scan(&l1, &l2, &l3)

	if l1 < l2+l3 && l2 < l1+l3 && l3 < l1+l3 {
		if l1 == l2 && l2 == l3 {
			fmt.Printf("Equilatero\n")
		} else if l1 == l2 || l1 == l3 || l2 == l3 {
			fmt.Printf("Isosceles\n")
		} else {
			fmt.Printf("Escaleno\n")
		}
	} else {
		fmt.Printf("Triangulo nao existe\n")
	}
}
