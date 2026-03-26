package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Printf("Digite dois números: ")
	fmt.Scan(&x, &y)

	if x%y == 0 {
		fmt.Printf("%d é divisível por %d\n", x, y)
	} else {
		fmt.Printf("%d não é divisível por %d\n", x, y)
	}
}
