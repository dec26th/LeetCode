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

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func TestGo(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	fmt.Println(<- c)
	fmt.Println(<- c)
}
