package maxProduct

import "math"

// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

func maxProduct(nums []int) int {
	max := math.MinInt64
	tempMax, tempMin := 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			tempMax, tempMin = tempMin, tempMax
		}

		tempMax = Max(nums[i], tempMax*nums[i])
		tempMin = Min(nums[i], tempMin*nums[i])
		max = Max(max, tempMax)
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
