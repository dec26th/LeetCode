package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	result = make([][]int, 0)

	searchPath(root, targetSum, []int{})
	return result
}

func searchPath(root *TreeNode, targetSum int, record []int) {
	if root == nil {
		return
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		record = append(record, root.Val)
		result = append(result, record)
		return
	}

	if root.Left != nil {
		temp := make([]int, len(record))
		copy(temp, record)
		temp = append(temp, root.Val)
		searchPath(root.Left, targetSum-root.Val, temp)
	}

	if root.Right != nil {
		temp := make([]int, len(record))
		copy(temp, record)
		temp = append(temp, root.Val)
		searchPath(root.Right, targetSum-root.Val, temp)
	}
}
