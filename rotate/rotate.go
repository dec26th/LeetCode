package rotate

func rotate(matrix [][]int) {
	max := len(matrix) / 2

	lengthOfMatrix := len(matrix) - 1
	for num := 0; num < max; num++ {

		for i := num; i < len(matrix) - 1 - num; i++ {
			a, b, c, d := matrix[num][i], matrix[i][lengthOfMatrix - num], matrix[lengthOfMatrix-num][lengthOfMatrix-i], matrix[lengthOfMatrix-i][num]
			matrix[num][i], matrix[i][lengthOfMatrix - num], matrix[lengthOfMatrix-num][lengthOfMatrix-i], matrix[lengthOfMatrix-i][num] = d, a, b, c
		}
	}
}
