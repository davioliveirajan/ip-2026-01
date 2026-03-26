package main

import (
	"fmt"
)

func main() {
	var x, y, z, maior, menor, inter int
	fmt.Printf("Digite 3 números: ")
	fmt.Scan(&x, &y, &z)

	if x > y && x > z {
		maior = x
	} else if y > x && y > z {
		maior = y
	} else {
		maior = z
	}
	if x < y && x < z {
		menor = x
	} else if y < x && y < z {
		menor = y
	} else {
		menor = z
	}
	if menor == x && maior == y {
		inter = z
	} else if menor == y && maior == z {
		inter = x
	} else {
		inter = y
	}

	fmt.Printf("%d, %d, %d\n", menor, inter, maior)
}
