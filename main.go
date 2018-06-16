package main

import (
	"fmt"

	"./algorytm"
)

func main() {
	array := []int{11, 7, 10, 5, 6}
	x := 1
	array = algorytm.QuickSort(array)
	result := algorytm.BinarySearch(array, x)

	if result != -1 {
		fmt.Println("x =", array[int(result)], "index = ", result)
	} else {
		fmt.Println("not found")
	}
}
