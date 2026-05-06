package main

import f"fmt"

func main() {
	var num[] int
	var div[] int
	var n int
	f.Printf("Digite os numeradores: ")
	for i := 0; i < 10; i++ {
		f.Scan(&n)
		num = append(num, n)
	}
	
	f.Printf("Digite os divisores: ")
	for i := 0; i < 5; i++ {
		f.Scan(&n)
		div = append(div, n)
	}		
	
	for _, l := range num {
	f.Printf("Número %d:\n", l)

	for i, v := range div {
		if v != 0 && l%v == 0 {
			f.Printf("Divisível por %d na posição %d\n", v, i)
		}
	}
}

}
 