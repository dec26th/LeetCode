package min_depth

type ListNode struct {
	 Val		int
	 Left		*ListNode
	 Right		*ListNode
}

var (
	visited map[*ListNode]bool
	queue []*ListNode
)

func minDepth(root *ListNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	queue = append(queue, root)

	for len(queue) != 0 {

		var indexToPoll int
		lenOfQueue := len(queue)
		for indexToPoll < lenOfQueue - 1  {
			tempNode := queue[0]
			queue := queue[1:]
			if tempNode.Left == nil && tempNode.Right == nil {
				return depth
			}

			if tempNode.Left != nil {
				queue = append(queue, tempNode.Left)
			}
			if tempNode.Right != nil {
				queue = append(queue, tempNode.Right)
			}

			indexToPoll ++
		}
		depth ++
	}
	return depth
}
