package copy_random_list

import (
	"fmt"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	node := new(Node)
	node2 := new(Node)

	node.Random = node2
	node.Next = node2
	node.Val = 1

	node2.Random = node2
	node2.Val = 2

	result := copyRandomListSplitCombine(node)
	for result != nil {
		fmt.Println(result.Val)
		fmt.Println(result.Random.Val)
		result = result.Next
	}
}

func TestCopyRandomList2(t *testing.T) {
	node1 := new(Node)
	node2 := new(Node)
	node3 := new(Node)
	node4 := new(Node)
	node5 := new(Node)

	node1.Val = 7
	node2.Val = 13
	node3.Val = 11
	node4.Val = 10
	node5.Val = 1

	node1.Next = node2

	node2.Next = node3
	node2.Random = node1

	node3.Next = node4
	node3.Random = node5

	node4.Next = node5
	node4.Random = node3

	node5.Random = node1

	result := copyRandomListSplitCombine(node1)
	for result != nil {
		fmt.Println(result.Val)
		fmt.Println(result.Random.Val)
		result = result.Next
	}
}
