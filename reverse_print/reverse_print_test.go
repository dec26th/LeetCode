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
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
	close(c)
	close(c)
}

func TestInterface(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	a <- 1
	b <- 2
	c <- 3
	select {

	case i := <-a:
		fmt.Println(i)
	case i := <-b:
		fmt.Println(i)
	case i := <-c:
		fmt.Println(i)
	}
}
