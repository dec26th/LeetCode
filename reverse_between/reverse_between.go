package reverse_between

type ListNode struct {
	Val		int
	Next	*ListNode
}

// 反转m-n之间的链表
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	count := 1

	for head != nil {
		if count >= m && count <= n {

		}

		if count > n {
			break
		}
	}

	return head
}
