package main

import "fmt"

func numIslands(grid [][]byte) int {
	var out int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				out++
				fill(grid, i, j)
			}
		}
	}
	return out
}

func fill(grid [][]byte, i int, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '0'
	fill(grid, i-1, j)
	fill(grid, i+1, j)
	fill(grid, i, j-1)
	fill(grid, i, j+1)
}

func main() {
	fmt.Println(numIslands([][]byte{{1, 1, 1}, {1, 1, 1}}))
}
