package algorytm

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
