package reverse_between

import "testing"

func TestReverseBetween(t *testing.T) {
	test := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	reverseBetween(test, 1, 2)
}
