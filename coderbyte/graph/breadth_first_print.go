package graph

func BreadthFirstPrint(graph map[string][]string, source string) []string {
	queue := []string{source}
	ret := []string{}

	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]

		ret = append(ret, c)
		for _, neighbor := range graph[c] {
			queue = append(queue, neighbor)
		}
	}
	return ret
}
