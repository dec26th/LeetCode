package detect_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for slow != nil && fast != nil {
		if slow != fast {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			return slow
		}
	}

	return nil
}
