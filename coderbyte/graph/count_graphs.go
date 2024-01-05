package graph

func CountGraphs(graph map[string][]string) int {
	count := 0
	visited := make(map[string]bool)
	for node := range graph {
		if _, ok := visited[node]; ok {
			continue
		}
		if ExploreGraphDFPRec(graph, node, visited) {
			count++
		}
	}
	return count
}

func ExploreGraphDFPRec(graph map[string][]string, node string, visited map[string]bool) bool {
	if _, ok := visited[node]; ok {
		return false
	}

	visited[node] = true

	for _, neighbor := range graph[node] {
		ExploreGraphDFPRec(graph, neighbor, visited)
	}

	return true
}
