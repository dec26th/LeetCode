package num_islands

var row, line int

func numIslands(grid [][]byte) int {
	row = len(grid)
	line = len(grid[0])
	if line == 1 && row == 1 {
		return int(grid[0][0])
	}
	numOfIsland := 0

	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if grid[i][j] == byte(1) {
				numOfIsland += 1
				findRelatedFriend(i, j, &grid)
			}
		}
	}
	return numOfIsland
}

func findRelatedFriend(x int, y int, grid *[][]byte) {
	if x == row || x < 0 {
		return
	}
	if y == line || y < 0 {
		return
	}
	if (*grid)[x][y] == 1 {
		(*grid)[x][y] = 0
		findRelatedFriend(x, y+1, grid)
		findRelatedFriend(x, y-1, grid)
		findRelatedFriend(x+1, y, grid)
		findRelatedFriend(x-1, y, grid)
	}
	return
}
