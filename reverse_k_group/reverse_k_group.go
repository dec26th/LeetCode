package reverse_k_group

type ListNode struct {
	Val		int
	Next	*ListNode
}
//todo 每k个反转一次
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}

	pre := dummyNode
	end := dummyNode

	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next // 记录起始点
		next := end.Next  // 记录下一个起始点
		end.Next = nil
		pre.Next = reverse(start)
		start.Next = next
		pre = start
		end = pre
	}

	return dummyNode.Next
}


func reverse(start *ListNode) *ListNode {
	var pre *ListNode
	curr := start
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}

	return pre
}
