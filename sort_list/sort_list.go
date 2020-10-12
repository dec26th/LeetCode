package sort_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	quickSort(head, nil)
	return head
}

func quickSort(head, end *ListNode) {
	if head == end || head.Next == end {
		return
	}

	pivot := head.Val
	cur, index := head.Next, head
	for cur != end {
		if cur.Val < pivot {
			index = index.Next
			cur.Val, index.Val = index.Val, cur.Val
		}
		cur = cur.Next
	}

	head.Val, index.Val = index.Val, head.Val
	quickSort(head, index)
	quickSort(index.Next, end)
}
