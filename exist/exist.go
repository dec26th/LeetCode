package exist

// 剑指 Offer 12. 矩阵中的路径

//请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，
//每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
//例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。
//
//[["a","b","c","e"],
//["s","f","c","s"],
//["a","d","e","e"]]
//
//但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

var (
	lenOfBoard  int
	lineOfBoard int
)

func exist(board [][]byte, word string) bool {
	byteWord := []byte(word)
	lenOfBoard = len(board)
	lineOfBoard = len(board[0])
	if len(byteWord) > lenOfBoard*lineOfBoard {
		return false
	}

	for i := 0; i < lenOfBoard; i++ {
		for j := 0; j < lineOfBoard; j++ {
			if board[i][j] == byteWord[0] {
				mark := buildMark(lenOfBoard, lineOfBoard)
				mark[i][j] = 1
				if bfs(board, byteWord[1:], i, j, mark) {
					return true
				}
			}
		}
	}
	return false
}

func buildMark(line, row int) [][]int {
	mark := make([][]int, line)
	for i := 0; i < line; i++ {
		mark[i] = make([]int, row)
	}
	return mark
}

func bfs(board [][]byte, word []byte, i, j int, mark [][]int) bool {
	if len(word) == 0 {
		return true
	}

	if i > 0 && board[i-1][j] == word[0] && mark[i-1][j] != 1 {
		mark[i-1][j] = 1
		if bfs(board, word[1:], i-1, j, mark) {
			return true
		}
		mark[i-1][j] = 0
	}

	if j > 0 && board[i][j-1] == word[0] && mark[i][j-1] != 1 {
		mark[i][j-1] = 1
		if bfs(board, word[1:], i, j-1, mark) {
			return true
		}
		mark[i][j-1] = 0
	}

	if i < lenOfBoard-1 && board[i+1][j] == word[0] && mark[i+1][j] != 1 {

		mark[i+1][j] = 1
		if bfs(board, word[1:], i+1, j, mark) {
			return true
		}
		mark[i+1][j] = 0
	}

	if j < lineOfBoard-1 && board[i][j+1] == word[0] && mark[i][j+1] != 1 {

		mark[i][j+1] = 1
		if bfs(board, word[1:], i, j+1, mark) {
			return true
		}
		mark[i][j+1] = 0
	}

	return false
}
