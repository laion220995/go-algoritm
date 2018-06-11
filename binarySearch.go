package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 5, 6}
	x := 0
	result := binarySearch(array, x)

	if result != -1 {
		fmt.Println(array[int(result)], result)
	} else {
		fmt.Println("not found")
	}
}

func binarySearch(array [5]int, x int) int {
	high := len(array) - 1
	low := 0

	for low <= high {
		mid := int((high + low) / 2)
		guess := array[mid]

		if guess == x {
			return mid
		} else if guess > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
