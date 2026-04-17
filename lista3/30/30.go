package main

import f"fmt"

func main() {
	var v, r float64

	for i := r; i <= 20; i += 0.5 {
		v = 4/3 * 3.14159 * r * r * r
		f.Printf("Volume da esfera de raio %.1f cm = %.2f cm³\n", r, v)
		r += 0.5
	}
}
