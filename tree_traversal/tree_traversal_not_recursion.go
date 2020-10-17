package tree_traversal

func preorderTraversalNotResursion(root *TreeNode) []int {
	stack := *new([]*TreeNode)
	result := *new([]int)

	for 0 < len(stack) || root != nil {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}

		index := len(stack) - 1
		root = stack[index]
		stack = stack[:index]
	}

	return result
}

func inorderTraversalNotResursion(root *TreeNode) []int {
	stack := *new([]*TreeNode)
	result := *new([]int)

	for 0 < len(stack) || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		index := len(stack) - 1
		root = stack[index]
		stack = stack[:index]

		result = append(result, root.Val)
		root = root.Right

	}

	return result
}

func postorderTraversalNotRecursion(root *TreeNode) []int {
	if root == nil {
		return []int {}
	}
	result := *new([]int)
	stack := *new([]*TreeNode)
	stack = append(stack, root)

	for len(stack) > 0 {
		root = stack[0]
		stack = stack[1:]
		result = append([]int {root.Val}, result...)

		if root.Left != nil {
			stack = append([]*TreeNode{root.Left}, stack...)
		}

		if root.Right != nil {
			stack = append([]*TreeNode{root.Right}, stack...)
		}
	}

	return  result
}