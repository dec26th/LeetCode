package rob_series

//198. 打家劫舍
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
//如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	result := -1

	dp := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
		result = max(dp[i], result)
	}
	return result
}

func robBetter(nums []int) int {
	lenOfNums := len(nums)

	var dpI1, dpI2, dpI int
	for i := lenOfNums - 1; i >= 0; i-- {
		dpI = max(dpI1, nums[i]+dpI2)
		dpI2 = dpI1
		dpI1 = dpI
	}
	return dpI
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 213. 打家劫舍 II

//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
//同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。

func robTwo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	return max(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}

func robRange(nums []int, start, end int) int {
	var dpI1, dpI2, dpI int

	for i := end; i >= start; i-- {
		dpI = max(dpI1, dpI2+nums[i])
		dpI2 = dpI1
		dpI1 = dpI
	}
	return dpI
}

// 337. 打家劫舍 III

//在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。
//除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
//如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
//
//计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func robThree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	takeIt := root.Val
	if root.Right != nil {
		takeIt += robThree(root.Right.Right) + robThree(root.Right.Left)
	}

	if root.Left != nil {
		takeIt += robThree(root.Left.Right) + robThree(root.Left.Left)
	}
	doNotTake := robThree(root.Left) + robThree(root.Right)

	return max(takeIt, doNotTake)
}
