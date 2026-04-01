package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Printf("A média é %d\n", media(x, y, z))
}

func media(x, y, z int) int {
	return (x + y + z) / 3
}