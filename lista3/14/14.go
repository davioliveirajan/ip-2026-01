package main

import f"fmt"

func main() {
	var x, y, z int
	f.Printf("Digite dois números: ")
	f.Scan(&x, &y)

	if y > x {
		z = y
		y = x
		x = z
	}

	for y <= x {
		primo := 1

		if y < 2 {
			primo = 0
		} else {
			i := 2
			for i < y {
				if y%i == 0 {
					primo = 0
					break
				}
				i++
			}
		}

		if primo == 1 {
			f.Printf("%d\n", y)
		}

		y++
	}
}