package trap

import "math"

//todo 接雨水
func trap(height []int) int {
	var totalDrop int
	var hasLow bool
	highest := getHighest(height)
	lenOfHeight := len(height)

	for i := 0; i < highest; i++ {
		temp := 0
		hasLow = false
		for j := 0; j < lenOfHeight; j++ {
			if hasLow && height[j] < i+1 {
				temp++
			}

			if height[j] >= i+1 {
				totalDrop += temp
				temp = 0
				hasLow = true
			}
		}
	}

	return totalDrop
}

func getHighest(height []int) (max int) {
	for _, v := range height {
		if max < v {
			max = v
		}
	}
	return max
}

func trapDP(height []int) int {
	var totalDrop int
	lenOfHeight := len(height)
	maxRight := make([]int, lenOfHeight)
	maxLeft := make([]int, lenOfHeight)

	for i := lenOfHeight - 2; i >= 0; i-- {
		maxRight[i] = int(math.Max(float64(maxRight[i+1]), float64(height[i+1])))
	}

	for i := 1; i < lenOfHeight-1; i++ {
		maxLeft[i] = int(math.Max(float64(maxLeft[i-1]), float64(height[i-1])))
	}

	for i := 1; i < lenOfHeight; i++ {
		min := int(math.Min(float64(maxLeft[i]), float64(maxRight[i])))
		if min > height[i] {
			totalDrop += min - height[i]
		}
	}
	return totalDrop
}
