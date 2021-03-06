package main

import (
	"bufio"
	"fmt"
	"os"

	"./algorytm"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	checkBinarySearch()

	checkBreadthFirstSearch()

	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func checkBinarySearch() {
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

func checkBreadthFirstSearch() {
	graph := make(map[string][]string)
	graph["Bob"] = []string{"Clem", "Alice"}
	graph["Clem"] = []string{"Alice", "Rob"}
	graph["Alice"] = []string{"Artur"}
	graph["Rob"] = []string{}
	graph["Artur"] = []string{"Bob", "Rob", "Alice"}

	fmt.Println(algorytm.BreadthFirstSearch(graph, "A", checkedFunc))
}

// checked condition for BreadthFirstSearch
func checkedFunc(value *string) bool {
	if len(*value) >= 5 {
		return true
	}
	return false
}
