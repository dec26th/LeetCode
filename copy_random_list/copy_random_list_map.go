package copy_random_list

func copyRandomListMap(head *Node) *Node {
	if head == nil {
		return nil
	}

	NodeMap := make(map[*Node]*Node)

	cur := head
	for head != nil {
		temp := new(Node)
		temp.Val = head.Val
		NodeMap[head] = temp

		head = head.Next
	}

	result := NodeMap[cur]
	temp := result
	for cur != nil {
		result.Next = NodeMap[cur.Next]
		result.Random = NodeMap[cur.Random]
		cur = cur.Next
		result = result.Next
	}

	return temp
}
