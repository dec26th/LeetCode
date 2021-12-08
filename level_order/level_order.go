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


func levelOrderToday(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{
	}
	layer := []*TreeNode{root}
	for len(layer) != 0 {

		newLayer := make([]*TreeNode, 0)
		thisLayer := make([]int, len(layer))
		for i := 0; i < len(layer); i++ {
			if layer[i].Left != nil {
				newLayer = append(newLayer, layer[i].Left)
			}
			if layer[i].Right != nil {
				newLayer = append(newLayer, layer[i].Right)
			}
			thisLayer[i] = layer[i].Val
 		}
		 result = append(result, thisLayer)
		 layer = newLayer
	}


	return result
}