package max_area_of_island

import "math"

var line, coluom int

func maxAreaOfIsland(grid [][]int) int {
	line = len(grid)
	coluom = len(grid[0])
	maxArea := math.MinInt32

	for i := 0; i < line; i++ {
		for j := 0; j < coluom; j++ {
			if grid[i][j] == 1 {
				area := getArea(i, j, &grid)
				maxArea = int(math.Max(float64(area), float64(maxArea)))
			}
		}
	}
	return maxArea
}

func getArea(x int, y int, grid *[][]int) int {
	if x == line || x < 0 {
		return 0
	}
	if y == coluom || y < 0 {
		return 0
	}

	if (*grid)[x][y] == 1 {
		(*grid)[x][y] = 0
		return 1 + getArea(x, y + 1, grid) + getArea(x + 1, y, grid) + getArea(x - 1, y, grid) + getArea(x, y - 1, grid)
	}
	return 0
}

