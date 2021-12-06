package max_area

func maxArea(height []int) int {

	result := 0
	for i := 0; i < len(height); i++ {

		result = max(result, min(height[i], height[len(height) - 1]) * (len(height) - i - 1))
		maxRight := height[len(height) - 1]
		for j := len(height) - 2; j >= 0 && j > i; j-- {
			if height[j] > maxRight {
				result = max(result, min(height[i], height[j]) * (j - i))
				maxRight = height[j]
			}
		}
	}

	return result
}


func min (x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func maxAreaBetter(height []int) int {
	left, right := 0, len(height) - 1

	result := 0
	for left < right {
		if height[left] < height[right] {
			result = max(result, height[left] * (right - left))
			left ++
		} else {
			result = max(result, height[right] * (right - left))
			right --
		}
	}
	return result
}