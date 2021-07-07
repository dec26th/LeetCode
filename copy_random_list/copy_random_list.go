package copy_random_list

//请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
//还有一个 random 指针指向链表中的任意节点或者 null。

type Node struct {
	Val int
	Next *Node
	Random *Node
}

var length = 1
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	record := head

	node := new(Node)
	result := new(Node)
	result.Next = node
	
	for head.Next != nil {
		node.Val = head.Val
		node.Next = new(Node)

		node = node.Next
		head = head.Next
		length ++
	}
	node.Val = head.Val

	target := result.Next
	for i := 0; i < length; i++ {
		random := record.Random
		target.Random = getRandomNode(result.Next, indexOfNode(random))

		record = record.Next
		target = target.Next
	}
	return result.Next
}

func indexOfNode(node *Node) int {
	index := 0

	for node != nil {
		node = node.Next
		index ++
	}
	return length - index
}

func getRandomNode(head *Node, index int) *Node {
	for i := 0; i < index; i++ {
		head = head.Next
	}

	return head
}