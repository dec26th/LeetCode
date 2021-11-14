package max_depth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return getDepth(root, 0)
}

func getDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	return max(getDepth(root.Left, depth+1), getDepth(root.Right, depth+1))
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
