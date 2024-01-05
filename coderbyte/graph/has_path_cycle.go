package graph

func BuildGraph(edges [][]string) map[string][]string {
	graph := make(map[string][]string)

	for _, row := range edges {
		a, b := row[0], row[1]
		if _, ok := graph[a]; !ok {
			graph[a] = []string{}
		}
		if _, ok := graph[b]; !ok {
			graph[b] = []string{}
		}
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	return graph
}

func HasPathCycleRec(graph map[string][]string, src, dst string) bool {
	visited := make(map[string]bool, 0)
	hasPath := hasPathCycleRecInternal(graph, src, dst, visited)
	return hasPath
}

func hasPathCycleRecInternal(graph map[string][]string, src, dst string, visited map[string]bool) bool {
	if _, ok := visited[src]; ok {
		/* Already visited before. No reason to do it once again */
		return false
	}

	if src == dst {
		return true
	}

	visited[src] = true

	for _, neighbor := range graph[src] {
		if hasPathCycleRecInternal(graph, neighbor, dst, visited) {
			return true
		}
	}

	return false
}
