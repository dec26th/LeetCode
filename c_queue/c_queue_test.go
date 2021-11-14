package c_queue

import (
	"fmt"
	"testing"
)

func TestCQueue_AppendTail(t *testing.T) {
	CQueue := Constructor()
	CQueue.AppendTail(1)
	CQueue.AppendTail(8)
	CQueue.AppendTail(20)
	CQueue.AppendTail(1)
	CQueue.AppendTail(11)
	CQueue.AppendTail(2)

	fmt.Println(CQueue.DeleteHead())
	fmt.Println(CQueue.DeleteHead())
	fmt.Println(CQueue.DeleteHead())
	fmt.Println(CQueue.DeleteHead())
	fmt.Println(CQueue.DeleteHead())
	fmt.Println(CQueue.DeleteHead())
}
