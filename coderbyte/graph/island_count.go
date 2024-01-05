package graph

import "fmt"

func IslandCount(grid [][]string) int {
	visited := make(map[string]bool)
	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if Explore(grid, r, c, visited) == true {
				count++
			}
		}
	}
	return count
}

func Explore(grid [][]string, r, c int, visited map[string]bool) bool {
	rowInbounds := 0 <= r && r < len(grid)
	colInbounds := 0 <= c && c < len(grid[0])

	if !rowInbounds || !colInbounds {
		return false
	}
	key := fmt.Sprintf("%d.%d", r, c)

	if _, ok := visited[key]; ok {
		return false
	}

	visited[key] = true

	if grid[r][c] == "W" {
		return false
	}

	Explore(grid, r-1, c, visited) // Up
	Explore(grid, r+1, c, visited) // Down
	Explore(grid, r, c-1, visited) // Left
	Explore(grid, r, c+1, visited) // Right

	return true
}
