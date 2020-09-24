package tree_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := new([]int)
	dfsIn(root, result)
	return *result
}

func dfsIn(root *TreeNode, result *[]int) {
	if root != nil {
		dfsIn(root.Left, result)
		(*result) = append(*result, root.Val)
		dfsIn(root.Right, result)
	}
}

func dfsPre(root *TreeNode, result *[]int) {
	if root != nil {
		(*result) = append(*result, root.Val)
		dfsPre(root.Left, result)
		dfsPre(root.Right, result)
	}
}

func preorderTraversal(root *TreeNode) []int {
	result := new([]int)
	dfsPre(root, result)
	return *result
}

func postoderTraversal(root *TreeNode) []int {
	result := new([]int)
	dfsPost(root, result)
	return *result
}

func dfsPost(root *TreeNode, result *[]int) {
	if root != nil {
		dfsPost(root.Left, result)
		dfsPost(root.Right, result)
		(*result) = append(*result, root.Val)
	}
}
