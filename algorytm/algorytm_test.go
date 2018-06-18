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

func testcheckedFunc1(value *string) bool {
	if len(*value) >= 5 {
		return true
	}
	return false
}
func testcheckedFunc2(value *string) bool {
	if len(*value) >= 6 {
		return true
	}
	return false
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
			t.Error("Compared array len's is not equal", len(resultArray), len(pair.resultArray))
		}

		for i := range resultArray {
			if resultArray[i] != pair.resultArray[i] {
				t.Error("Array's elements is not equal", resultArray[i], pair.resultArray[i])
			}
		}
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	graph := make(map[string][]string)
	graph["Bob"] = []string{"Clem", "Alice"}
	graph["Clem"] = []string{"Alice", "Rob"}
	graph["Alice"] = []string{"Artur"}
	graph["Rob"] = []string{}
	graph["Artur"] = []string{"Bob", "Rob", "Alice"}

	// function has an element passing the test
	result := BreadthFirstSearch(graph, "Bob", testcheckedFunc1)

	if result != true {
		t.Error("BreadthFirstSearch return true, need false (testcheckedFunc1)")
	}

	// function doesn't have an element passing the test
	result = BreadthFirstSearch(graph, "Bob", testcheckedFunc2)

	if result != false {
		t.Error("BreadthFirstSearch return false, need true (testcheckedFunc2)")
	}

	// function tests an element that is missing in the graph
	result = BreadthFirstSearch(graph, "Derek", testcheckedFunc2)

	if result != false {
		t.Error("BreadthFirstSearch return false, need true (testcheckedFunc2)")
	}
}
