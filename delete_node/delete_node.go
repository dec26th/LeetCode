package delete_node

// 剑指 Offer 18. 删除链表的节点]

// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
//
//返回删除后的链表的头节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dum := &ListNode{
		Next: head,
	}
	cur := dum

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			return dum.Next
		}
		cur = cur.Next
	}
	return dum.Next
}
