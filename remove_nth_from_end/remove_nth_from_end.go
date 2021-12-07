package remove_nth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeWithBefore struct {
	Next   *ListNodeWithBefore
	Before *ListNodeWithBefore
	Node   *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headBefore := &ListNode{Next: head}
	cur := head

	begin := new(ListNodeWithBefore)
	begin.Node = head
	begin.Before = &ListNodeWithBefore{Next: begin, Node: headBefore}

	for cur != nil {
		begin.Next = new(ListNodeWithBefore)
		begin.Next.Before = begin
		begin.Next.Node = cur.Next

		begin = begin.Next
		cur = cur.Next
	}

	for i := 0; i <= n; i++ {
		begin = begin.Before
	}

	if begin != nil && begin.Node != nil && begin.Node.Next != nil {
		begin.Node.Next = begin.Node.Next.Next
	} else {
		head = nil
	}
	return headBefore.Next
}
