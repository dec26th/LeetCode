package solve_n_queens

var result [][]string
var num int

func solveNQueens(n int) [][] string {
	num = n
	board := new([]string)
	for i := 0; i < n; i++ {
		*board = append(*board, "....")
	}
	backTrack(board, n)
	return result
}

func backTrack(board *[]string, raw int) {
	if raw == len(*board) {
		result = append(result, *board)
		return
	}

	for col := 0; col < num; col++ {
		if !isVaild(board, raw, col) {
			continue
		}

	}

}

func isVaild(board *[]string, row, col int) bool {
	for i := 0; i < num; i++ {
		if (*board)[i][col] == 'Q' {
			return false
		}
	}

	i := col + 1
	for j := row - 1; j < num && i >= 0; j++ {
		if (*board)[i][j] == 'Q' {
			return false
		}
		i--
	}

	i = row - 1
	j := col - 1
	for i >=0 && j >= 0 {
		if (*board)[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	return true
}