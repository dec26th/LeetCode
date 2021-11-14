package reverse_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dump := make([]int, 0)

	for head != nil {
		dump = append(dump, head.Val)
		head = head.Next
	}

	head = new(ListNode)
	result := head
	for i := len(dump) - 1; i >= 0; i-- {
		result.Val = dump[i]
		if i == 0 {
			break
		}
		result.Next = new(ListNode)
		result = result.Next
	}
	result.Next = nil
	return head
}

func reverseListTowPointer(head *ListNode) *ListNode {
	pre := &ListNode{}
	cur := head

	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reverse_list_two(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return temp
}
