package algorytm

import (
	"testing"
)

type testbinary struct {
	array  []int
	x      int
	result int
}

type testquicksort struct {
	array       []int
	resultArray []int
}

var binaryTest = []testbinary{
	{[]int{1, 2, 3, 4, 5}, 5, 4},
	{[]int{1, 2, 3, 4, 5}, 6, -1},
}

var quickSortTest = []testquicksort{
	{[]int{1}, []int{1}},
	{[]int{1, 4, 3, 2, 6}, []int{1, 2, 3, 4, 6}},
	{[]int{2, 1, 2, 1, 2}, []int{1, 1, 2, 2, 2}},
}

func TestBinarySearch(t *testing.T) {
	for _, pair := range binaryTest {
		result := BinarySearch(pair.array, pair.x)

		if result != pair.result {
			t.Fatalf("Expected %v -> got %v", pair.result, result)
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, pair := range quickSortTest {
		resultArray := QuickSort(pair.array)

		if resultArray == nil {
			t.Fatalf("Expected %v -> got %v", pair.resultArray, resultArray)
		}

		if len(resultArray) != len(pair.resultArray) {
			t.Error("Compare array len is not equal", len(resultArray), len(pair.resultArray))
		}

		for i := range resultArray {
			if resultArray[i] != pair.resultArray[i] {
				t.Error("Array element is not equal", resultArray[i], pair.resultArray[i])
			}
		}
	}
}
