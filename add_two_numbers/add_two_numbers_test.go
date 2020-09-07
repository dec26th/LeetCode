package add_two_numbers

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := & ListNode{
		Val:  3,
		Next: &ListNode{
			Val: 7,
		},
	}
	l2 :=  & ListNode{
		Val:  9,
		Next: &ListNode{
			Val: 2,
		},
	}

	fmt.Println(addTwoNumbers(l1, l2))
}
