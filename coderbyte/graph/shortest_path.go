package graph

import "math"

type Cost struct {
	Node     string
	Distance int
}

func TranslateToMap(edges [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
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

func FindShortestPath(edges [][]string, start, end string) int {
	shortest := math.MaxInt
	graph := TranslateToMap(edges)
	visited := make(map[string]bool)

	startCost := Cost{
		Distance: 0,
		Node:     start,
	}

	queue := []Cost{startCost}

	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]

		if _, ok := visited[c.Node]; ok {
			continue
		}

		visited[c.Node] = true

		if c.Node == end {
			shortest = int(math.Min(float64(shortest), float64(c.Distance)))
			continue
		}

		for _, neighbor := range graph[c.Node] {
			nCost := Cost{
				Node:     neighbor,
				Distance: c.Distance + 1,
			}
			queue = append(queue, nCost)
		}
	}

	return shortest
}
