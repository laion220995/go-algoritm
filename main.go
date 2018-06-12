package main

import (
	"fmt"

	"./algorytm"
)

func main() {
	array := []int{1, 2, 3, 5, 6}
	x := 1
	result := algorytm.BinarySearch(array, x)

	if result != -1 {
		fmt.Println("x =", array[int(result)], "index = ", result)
	} else {
		fmt.Println("not found")
	}
}
