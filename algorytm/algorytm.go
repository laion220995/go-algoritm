package algorytm

// BinarySearch - algorytm for binary search
func BinarySearch(array []int, x int) int {
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

// QuickSort - algorytm for quick sort using slice
func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	pivotIndex := len(array) / 2
	pivot := array[pivotIndex]
	var less []int
	var greater []int

	for i, item := range array {
		if i == pivotIndex {
			continue
		}

		if item <= pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}

	less = append(QuickSort(less), pivot)
	greater = QuickSort(greater)

	return append(less, greater...)
}
