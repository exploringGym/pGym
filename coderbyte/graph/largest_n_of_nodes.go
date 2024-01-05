package graph

import "math"

func FindLargestNumberOfNodes(graph map[string][]string) int {
	largest := 0
	visited := make(map[string]bool)

	for node := range graph {
		if _, ok := visited[node]; ok {
			continue
		}
		cSize := ExploreSizeDFS(graph, node, visited)
		largest = int(math.Max(float64(cSize), float64(largest)))
	}

	return largest
}

func ExploreSizeDFS(graph map[string][]string, node string, visited map[string]bool) int {
	if _, ok := visited[node]; ok {
		return 0
	}

	visited[node] = true

	count := 1

	for _, neighbor := range graph[node] {
		count += ExploreSizeDFS(graph, neighbor, visited)
	}

	return count
}
