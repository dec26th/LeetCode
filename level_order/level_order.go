package level_order

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	var order []*TreeNode

	order = append(order, root)

	for len(order) != 0 {
		mostLeft := order[0]
		order = order[1:]
		result = append(result, mostLeft.Val)

		if mostLeft.Left != nil {
			order = append(order, mostLeft.Left)
		}
		if mostLeft.Right != nil {
			order = append(order, mostLeft.Right)
		}
	}

	return result
}
