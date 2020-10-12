package merge_two_lists

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	valueList := *new([]int)

	for l1 != nil {
		valueList = append(valueList, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		valueList = append(valueList, l2.Val)
		l2 = l2.Next
	}

	sort.Ints(valueList)
	dumpNode := &ListNode{}
	head := dumpNode
	for _, v := range valueList {
		head.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		head = head.Next
	}
	return dumpNode.Next
}

func mergeTwoListsBetter(l1 *ListNode, l2 *ListNode) *ListNode {
	dumNode := &ListNode{}
	head := dumNode

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next

		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}

	return dumNode.Next
}
