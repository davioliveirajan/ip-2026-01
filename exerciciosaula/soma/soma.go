package main

import "fmt"

func main(){
	var x, y int
	fmt.Printf("Digite dois números: ")
	fmt.Scan(&x, &y)
	fmt.Printf("A soma é %d\n", soma(x, y))
}
func soma(x, y int) int {
	return x + y
}
