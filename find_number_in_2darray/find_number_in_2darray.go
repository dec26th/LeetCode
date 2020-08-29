package find_number_in_2darray

func findNumberIn2DArray(matrix [][]int, target int) bool {
	lenOfLine := len(matrix)
	if lenOfLine == 0 {
		return false
	}
	lenOfColumn := len(matrix[0])

	for i := 0; i < lenOfLine; i++ {
		for j := 0; j < lenOfColumn; j++ {
			if target == matrix[i][j] {
				return true
			}
		}
	}
	return false
}

func findNumberIn2DArrayLine(matrix [][]int, target int) bool {
	lenOfLine := len(matrix)
	if lenOfLine == 0 {
		return false
	}
	lenOfColumn := len(matrix[0])
	line, colum := 0, lenOfColumn-1
	for line < lenOfLine && colum >= 0 {
		if matrix[line][colum] == target {
			return true
		}
		if matrix[line][colum] > target {
			colum -= 1

		} else if matrix[line][colum] < target {
			line += 1

		}
	}
	return false
}
