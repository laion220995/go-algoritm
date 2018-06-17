package algorytm

import (
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

type checkedfunc func(value *string) bool

// BreadthFirstSearch - algorithm for traversing or
// searching tree or graph data structures
func BreadthFirstSearch(graph map[string][]string, name string, checkedFunc checkedfunc) bool {

	deque := deque.New()
	deque.PushRight(name)

	var searched []string
	for deque.Size() > 0 {

		person := deque.PopLeft()
		value, ok := person.(string)
		if ok {
			isChecked := false
			for _, searchedName := range searched {
				if searchedName == value {
					isChecked = true
				}
			}

			if !isChecked {
				if checkedFunc(&value) {
					return true
				} else if len(value) > 0 {
					for _, children := range graph[value] {
						deque.PushRight(children)
					}

					searched = append(searched, value)
				}
			}
		}
	}

	return false
}
