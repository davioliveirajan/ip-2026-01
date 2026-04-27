package main

import f"fmt"

func main() {
	l := [] int {32, 43, 2, 45, 3, 7, 37, 1}

	insertionsort(l)

	for i := 0; i < len(l); i++ {
		f.Printf("%d ", l[i])
	}
	f.Println()

}

func insertionsort(arr[] int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}