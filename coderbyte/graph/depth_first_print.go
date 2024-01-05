package graph

func DepthFirstPrint(graph map[string][]string, source string) []string {
	stack := make([]string, 0)
	stack = append(stack, source)

	ret := make([]string, 0)

	for len(stack) != 0 {
		l := len(stack) - 1
		c := stack[l]
		stack = stack[:l]
		ret = append(ret, c)
		for _, neighbor := range graph[c] {
			stack = append(stack, neighbor)
		}
	}
	return ret
}

func DepthFirstPrintRec(graph map[string][]string, source string) []string {
	ret := []string{source}

	node := graph[source]

	for _, neighbor := range node {
		ret = append(ret, DepthFirstPrint(graph, neighbor)...)
	}

	return ret
}
