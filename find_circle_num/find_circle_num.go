package find_circle_num

var row, line int

func findCircleNum(M [][]int) int {
	row = len(M)
	line = len(M[0])
	if line == 1 && row == 1 {
		return M[0][0]
	}
	numOfCircle := 0
	for i := 0; i < row; i++ {
		for j := 0; j < line; j++ {
			if M[i][j] == 1 {
				numOfCircle++
				findRelatedFriends(i, j, &M)
			}
		}
	}
	return numOfCircle
}

func findRelatedFriends(x int, y int, M *[][]int) {
	if x == row || x < 0 {
		return
	}
	if y == line || y < 0 {
		return
	}

	if (*M)[x][y] == 1 {
		(*M)[x][y] = 0
		for i := 0; i < row; i++ {
			findRelatedFriends(i, y, M)
		}
		for j := 0; j < line; j++ {
			findRelatedFriends(x, j, M)
		}
	}
}
