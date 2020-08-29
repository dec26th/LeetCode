package reverse_print

import (
	"fmt"
	"testing"
)

func TestReversePrint(t *testing.T) {
	var head ListNode
	head.Val = 1
	head.Next = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	fmt.Println(reversePrint(&head))
}
