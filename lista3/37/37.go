package main

import f"fmt"

func main() {
	var dec int
	var oct string
	f.Print("Digite um número em base 8: ")
	f.Scan(&oct)
	f.Sscanf(oct, "%o", &dec)

	f.Printf("%d\n", dec)
}