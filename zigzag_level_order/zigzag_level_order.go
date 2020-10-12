package zigzag_level_order

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	stack1 []*TreeNode
	stack2 []*TreeNode
	result [][]int
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result = *new([][]int)
	stack1 = *new([]*TreeNode)
	stack2 = *new([]*TreeNode)

	stack1 = append(stack1, root)
	for len(stack1) != 0 || len(stack2) != 0 {
		tempResult1 := *new([]int)
		tempResult2 := *new([]int)
		for len(stack1) != 0 {
			index := len(stack1) - 1
			tempNode := stack1[index]

			stack1 = stack1[:index]
			tempResult1 = append(tempResult1, tempNode.Val)
			if tempNode.Left != nil {
				stack2 = append(stack2, tempNode.Left)
			}
			if tempNode.Right != nil {
				stack2 = append(stack2, tempNode.Right)
			}
		}
		if len(tempResult1) != 0 {
			result = append(result, tempResult1)
		}

		for len(stack2) != 0 {
			index := len(stack2) - 1
			tempNode := stack2[index]

			stack2 = stack2[:index]
			tempResult2 = append(tempResult2, tempNode.Val)
			if tempNode.Right != nil {
				stack1 = append(stack1, tempNode.Right)
			}
			if tempNode.Left != nil {
				stack1 = append(stack1, tempNode.Left)
			}
		}
		if len(tempResult2) != 0 {
			result = append(result, tempResult2)
		}

	}

	return result
}
