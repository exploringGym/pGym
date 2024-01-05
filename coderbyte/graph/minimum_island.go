package graph

import (
	"fmt"
	"math"
)

func FindMinimumSizeIsland(grid [][]string) int {
	smallestIsland := math.MaxInt
	visited := make(map[string]bool)
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			islandSize := ExploreGridForIslandSizeDSFRec(grid, r, c, visited)
			if islandSize > 0 {
				smallestIsland = int(math.Min(float64(smallestIsland), float64(islandSize)))
			}
		}
	}
	if smallestIsland == math.MaxInt {
		smallestIsland = 0
	}
	return smallestIsland
}

func ExploreGridForIslandSizeDSFRec(grid [][]string, r, c int, visited map[string]bool) int {
	rowInbounds := r >= 0 && r < len(grid)
	colInbounds := c >= 0 && c < len(grid[0])

	if !rowInbounds || !colInbounds {
		return 0
	}

	key := fmt.Sprintf("%d.%d", r, c)
	if _, ok := visited[key]; ok {
		return 0
	}

	visited[key] = true

	if grid[r][c] == "W" {
		return 0
	}

	islandSize := 1
	islandSize += ExploreGridForIslandSizeDSFRec(grid, r-1, c, visited) // Up
	islandSize += ExploreGridForIslandSizeDSFRec(grid, r+1, c, visited) // Down
	islandSize += ExploreGridForIslandSizeDSFRec(grid, r, c-1, visited) // Left
	islandSize += ExploreGridForIslandSizeDSFRec(grid, r, c+1, visited) // Right

	return islandSize
}
