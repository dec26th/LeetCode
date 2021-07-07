package copy_random_list

func copyRandomListSplitCombine(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head

	for cur != nil {
		record := cur.Next

		cur.Next = new(Node)
		cur.Next.Val = cur.Val
		cur.Next.Next = record

		cur = record
	}

	cur = head
	nxt := cur.Next

	for nxt.Next != nil {
		if cur.Random == nil {
			nxt.Random = nil
		} else {
			nxt.Random = cur.Random.Next
		}
		cur = nxt.Next
		nxt = cur.Next
	}
	if cur.Random == nil {
		nxt.Random = nil
	} else {
		nxt.Random = cur.Random.Next
	}

	cur = head
	nxt = cur.Next
	result := nxt
	for nxt.Next != nil {
		cur.Next = nxt.Next

		nxt.Next = cur.Next.Next

		cur = cur.Next
		nxt = nxt.Next
	}
	cur.Next = nil

	return result
}
