package reverse_between

type ListNode struct {
	Val  int
	Next *ListNode
}

var left *ListNode

var stop bool

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	left = head
	stop = false
	recurseAndReverse(head, m, n)
	return head
}

func recurseAndReverse(right *ListNode, m int, n int) {
	if n == 1 {
		return
	}

	right = right.Next

	if m > 1 {
		left = left.Next
	}
	recurseAndReverse(right, m-1, n-1)

	if left == right || right.Next == left {
		stop = true
	}

	if !stop {
		t := left.Val
		left.Val = right.Val
		right.Val = t
		left = left.Next
	}
}

func reverseBetweenPointer(head *ListNode, m int, n int) *ListNode { // ✅
	dummyNode := &ListNode{
		Val: -1,
	}

	dummyNode.Next = head
	preNode := dummyNode
	for i := 1; i < m; i++ {
		preNode = preNode.Next
	}

	curNode := preNode.Next  // cur 表示当前节点， pre表示上一个节点， nxt表示下一个节点
	for i := m; i < n; i++ { // 每次将cur的next(也就是nxt)插到pre的后面
		nxtNode := curNode.Next
		curNode.Next = nxtNode.Next
		nxtNode.Next = preNode.Next
		preNode.Next = nxtNode
	}

	return dummyNode.Next
}
