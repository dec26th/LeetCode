package get_kth_from_end

import "testing"

func TestGetKthFromEnd(t *testing.T) {
	test := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	getKthFromEnd(test, 2)
}
