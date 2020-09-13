package tree_traversal

func preorderTraversalNotResursion(root *TreeNode) []int {
	stack := *new([] *TreeNode)
	result := *new([] int)

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
	stack := *new([] *TreeNode)
	result := *new([] int)

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
