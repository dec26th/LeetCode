package max_area_of_island

import (
	"math"
)

var row, line int

func maxAreaOfIsland(grid [][]int) int {
	row = len(grid)
	line = len(grid[0])
	if row == 1 && line == 1 {
		return grid[0][0]
	}
	maxArea := 0

	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if grid[i][j] == 1 {
				area := getArea(i, j, &grid)
				maxArea = int(math.Max(float64(area), float64(maxArea)))
			}
		}
	}
	return maxArea
}

func getArea(x int, y int, grid *[][]int) int {
	if x == row || x < 0 {
		return 0
	}
	if y == line || y < 0 {
		return 0
	}

	if (*grid)[x][y] == 1 {
		(*grid)[x][y] = 0
		return 1 + getArea(x, y+1, grid) + getArea(x+1, y, grid) + getArea(x-1, y, grid) + getArea(x, y-1, grid)
	}
	return 0
}
