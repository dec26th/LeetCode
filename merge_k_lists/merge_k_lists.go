package merge_k_lists

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 || (len(lists) == 1 && lists[0] == nil) {
		return nil
	}

	head := new(ListNode)
	body := head
	count := 0
	valueLists := new([]int)

	for _, v := range lists {
		if !appendValue(valueLists, v) {
			count++
		}
	}
	if count == len(lists) {
		return nil
	}

	sort.Ints(*valueLists)

	for k, v := range *valueLists {
		body.Val = v
		if k != len(*valueLists)-1 {
			body.Next = &ListNode{}
			body = body.Next
		}
	}
	body.Next = nil
	return head
}

func appendValue(valueLists *[]int, node *ListNode) bool {
	if node == nil {
		return false
	}

	for node != nil {
		*valueLists = append(*valueLists, node.Val)
		node = node.Next
	}
	return true
}
