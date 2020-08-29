package reverse_print

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	stack := make([]int, 0)

	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	lenOfStack := len(stack)
	result := make([]int, lenOfStack)
	for i := lenOfStack - 1; i >= 0; i-- {
		result[lenOfStack-i-1] = stack[i]
	}
	return result
}
