package spiral_order

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	result := make([]int, len(matrix)*len(matrix[0]))
	index := 0

	for {
		for i := left; i < right+1; i++ {
			result[index] = matrix[top][i]
			index++
		}
		top++
		if top > bottom {
			break
		}

		for i := top; i < bottom+1; i++ {
			result[index] = matrix[i][right]
			index++
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			result[index] = matrix[bottom][i]
			index++
		}
		bottom--
		if bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			result[index] = matrix[i][left]
			index++
		}
		left++
		if left > right {
			break
		}
	}
	return result
}
