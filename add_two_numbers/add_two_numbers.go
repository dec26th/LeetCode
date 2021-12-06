package add_two_numbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pass_num := 0
	count := 0
	head, body, lastBody := &ListNode{}, &ListNode{}, &ListNode{}

	for l1 != nil && l2 != nil {

		body.Next = &ListNode{}

		if count == 0 {
			count += 1
			head = body
		}

		body.Val = ((l1.Val + l2.Val) + pass_num) % 10
		pass_num = (l1.Val + l2.Val + pass_num) / 10
		lastBody = body
		body = body.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	lastBody.Next = nil
	body = lastBody

	for l1 != nil {
		body.Next = &ListNode{
			Val: (l1.Val + pass_num) % 10,
		}
		pass_num = (l1.Val + pass_num) / 10
		lastBody = body
		body = body.Next
		l1 = l1.Next
	}

	for l2 != nil {
		body.Next = &ListNode{
			Val: (l2.Val + pass_num) % 10,
		}
		pass_num = (l2.Val + pass_num) / 10
		lastBody = body
		body = body.Next
		l2 = l2.Next
	}

	if pass_num == 1 {
		body.Next = &ListNode{
			Val: pass_num,
		}
	}

	return head
}

func addTwoNumbersTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	cur := head
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		cur.Next = &ListNode{}
		cur = cur.Next

		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Val = carry % 10
		carry /= 10
	}

	return head.Next
}

func addTwoNumbersThree(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)

	cur := head
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		cur.Next = new(ListNode)
		cur = cur.Next

		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Val = carry % 10
		carry = carry / 10
	}

	return head.Next
}
