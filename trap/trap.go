package trap


//todo 接雨水
func trap(height []int) int {
	var totalDrop int
	highest := getHighest(height)
	lenOfHeight := len(height)

	for i := 0; i < highest; i++ {
		leftBarrier := 0
		rightBarrier := 0
		for j := 0; j < lenOfHeight; j++ {
			if height[j] > i && leftBarrier <= rightBarrier{
				leftBarrier = j
				continue
			}
			if height[j] > i && rightBarrier <= leftBarrier {
				totalDrop += j - leftBarrier - 1
				rightBarrier = j
				leftBarrier = j
				continue
			}
		}
	}

	return lenOfHeight
}

func getHighest(height []int) (max int) {
	for _, v := range height {
		if max < v {
			max = v
		}
	}
	return max
}
