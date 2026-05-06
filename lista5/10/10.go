package main

import f"fmt"

func main() {
	var n1, n2, fibon int
	var fibonacci[] int

	n1 = 0
	n2 = 1
	f.Printf("%d %d ", n1, n2)

	for i := 0; i < 48; i++ {
		fibon = n1 + n2
		n1 = n2
		n2 = fibon
		fibonacci = append(fibonacci, fibon)
		f.Printf("%d ", fibonacci[i])
	}

}