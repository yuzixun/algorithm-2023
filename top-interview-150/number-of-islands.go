package main

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				numIslandsHandle(grid, i, j)
			}
		}
	}
	return count
}

func numIslandsHandle(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	numIslandsHandle(grid, i+1, j)
	numIslandsHandle(grid, i-1, j)
	numIslandsHandle(grid, i, j+1)
	numIslandsHandle(grid, i, j-1)
}
