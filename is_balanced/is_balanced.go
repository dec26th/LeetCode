package is_balanced

type TreeNode struct {
	Val		int
	Left 	*TreeNode
	Right 	*TreeNode
}

func isBalanced(root *TreeNode) bool {
	return isChildrenBalanced(root)
}

func getDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth += 1
	return max(getDepth(root.Left, depth), getDepth(root.Right, depth))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isChildrenBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var depthLeft, depthRight int
	if root.Left != nil {
		depthLeft = getDepth(root.Left, 0)
	}
	if root.Right != nil {
		depthRight = getDepth(root.Right, 0)
	}

	if abs(depthLeft - depthRight) > 1 {
		return false
	}

	return isChildrenBalanced(root.Left) && isChildrenBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}