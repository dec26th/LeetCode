package kth_largest

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

var result int
var kth int

func kthLargest(root *TreeNode, k int) int {
	kth = k
	getKthLargest(root)
	return result
}

func getKthLargest(root *TreeNode) {
	if root == nil {
		return
	}
	getKthLargest(root.Right)
	if kth == 0 {
		return
	}
	kth --
	if kth == 0 {
		result = root.Val
		return
	}
	getKthLargest(root.Left)
}
