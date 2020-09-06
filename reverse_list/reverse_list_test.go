package reverse_list

import "testing"

func TestReverseList(t *testing.T) {
	test := &ListNode{
		Val:  0,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	reverseList(test)
}
