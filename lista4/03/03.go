package main 

import f"fmt"

func main() {
	n := []int{1, 6, 15, 11, 20}
    resultado := invert(n)
    f.Printf("Array invertido: %v\n", resultado)
}

func invert(arr[] int) []int {
	if len(arr) <= 1 {
		return arr
	}
	return append([]int{arr[len(arr)-1]}, invert(arr[:len(arr)-1])...)
}