package detect_cycle

// 142. 环形链表 II
// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

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
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}

	return nil
}
