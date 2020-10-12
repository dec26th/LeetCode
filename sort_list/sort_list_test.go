package sort_list

import "testing"

func TestSortList(t *testing.T) {
	test := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	sortList(test)
}
