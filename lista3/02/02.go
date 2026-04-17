package main

import f"fmt"

func main() {
	var soma, media int
	for x := 70; x >=50; x -= 2 {
		soma += x
	}
	media = soma / 11
	f.Printf("A média é: %d\n", media)
}