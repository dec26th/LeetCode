package reverse_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	dump := make([]int, 0)

	for head != nil {
		dump = append(dump, head.Val)
		head = head.Next
	}

	head = new(ListNode)
	result := head
	for i := len(dump) - 1; i >= 0; i-- {
		result.Val = dump[i]
		if i == 0 {
			break
		}
		result.Next = new(ListNode)
		result = result.Next
	}
	result.Next = nil
	return head
}

func reverseListTowPointer(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	count := 0
	cur := new(ListNode)
	pre := head

	for pre != nil {
		if count == 0 {
			temp := pre.Next
			pre.Next = nil
			cur = pre
			pre = temp
			count++
			continue
		}
		temp := pre.Next
		pre.Next = cur
		cur = pre
		pre = temp
		count++
	}
	return cur
}

func reverse_list_two(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return temp
}