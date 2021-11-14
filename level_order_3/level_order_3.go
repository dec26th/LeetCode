package level_order_3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	nodeList := []*TreeNode{
		root,
	}
	length := len(nodeList)
	left2Right, right2Left := true, false
	result := make([][]int, 0)

	for length != 0 {
		temp := make([]int, length)
		for i := 0; i < length; i++ {
			temp[i] = nodeList[i].Val
			if left2Right {
				if nodeList[length-1-i].Right != nil {
					nodeList = append(nodeList, nodeList[length-1-i].Right)
				}
				if nodeList[length-1-i].Left != nil {
					nodeList = append(nodeList, nodeList[length-1-i].Left)
				}
			} else if right2Left {
				if nodeList[length-1-i].Left != nil {
					nodeList = append(nodeList, nodeList[length-1-i].Left)
				}
				if nodeList[length-1-i].Right != nil {
					nodeList = append(nodeList, nodeList[length-1-i].Right)
				}
			}

		}
		result = append(result, temp)
		nodeList = nodeList[length:]
		left2Right = !left2Right
		right2Left = !right2Left
		length = len(nodeList)
	}
	return result
}
