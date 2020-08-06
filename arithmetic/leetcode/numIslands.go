package main

import "fmt"

func main() {
	grid := make([][]byte, 0)
	//grid = append(grid,
	//	[]byte{
	//		'1', '1', '0', '1', '0',
	//	},
	//	[]byte{
	//		'0', '1', '0', '1', '0',
	//	},
	//	[]byte{
	//		'1', '0', '0', '0', '0',
	//	},
	//	[]byte{
	//		'1', '1', '0', '1', '0',
	//	})
	grid = append(grid, []byte{'1'}, []byte{'1'})
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
				count++
			}
		}
	}
	return count
}

func dfs(grids [][]byte, i, j int) int {
	if i < 0 || i >= len(grids) || j < 0 || j >= len(grids[0]) {
		return 0
	}
	fmt.Println(i, j, len(grids))
	if grids[i][j] == '1' {
		grids[i][j] = 0
		return 1 + dfs(grids, i-1, j) + dfs(grids, i, j-1) + dfs(grids, i+1, j) + dfs(grids, i, j+1)
	}
	return 0
}
