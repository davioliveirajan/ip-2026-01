package main

import f"fmt"

func main() {
	n := []int{1, 6, 15, 11, 20}
	f.Printf("A soma dos valores é %d\n", soma(n))

}
 func soma(n[] int) int {
	if len(n) == 0 {
		return 0
	}
	return n[0] + soma(n[1:])
 }