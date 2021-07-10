package level_order_2



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	preLevelNode := []*TreeNode{
		root,
	}
	curLevelNode := make([]*TreeNode, 0)

	for len(preLevelNode) != 0 {
		temp := make([]int, 0)
		for i := 0; i < len(preLevelNode); i++ {
			temp = append(temp, preLevelNode[i].Val)

			if preLevelNode[i].Left != nil {
				curLevelNode = append(curLevelNode, preLevelNode[i].Left)
			}
			if preLevelNode[i].Right != nil {
				curLevelNode = append(curLevelNode, preLevelNode[i].Right)
			}
		}
		result = append(result, temp)

		preLevelNode = make([]*TreeNode, len(curLevelNode))
		copy(preLevelNode, curLevelNode)
		curLevelNode = make([]*TreeNode, 0)
	}

	return result
}