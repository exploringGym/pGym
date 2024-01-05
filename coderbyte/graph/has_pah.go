package graph

func hasPathDFS(graph map[string][]string, src, dst string) bool {
	stack := []string{src}
	for len(stack) != 0 {
		l := len(stack) - 1
		c := stack[l]

		if c == dst {
			return true
		}

		stack = stack[:l]
		for _, neighbor := range graph[c] {
			stack = append(stack, neighbor)
		}
	}
	return false
}

func hasPathDFSRec(graph map[string][]string, src, dst string) bool {
	if src == dst {
		return true
	}
	for _, neighbor := range graph[src] {
		if hasPathDFSRec(graph, neighbor, dst) == true {
			return true
		}
	}
	return false
}

func hasPathBST(graph map[string][]string, src, dst string) bool {
	queue := []string{src}
	for len(queue) != 0 {
		c := queue[0]
		if c == dst {
			return true
		}

		queue = queue[1:]
		for _, neighbor := range graph[c] {
			queue = append(queue, neighbor)
		}
	}
	return false
}
