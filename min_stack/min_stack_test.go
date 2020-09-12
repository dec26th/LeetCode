package min_stack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	test := Constructor()
	test.Push(2)
	test.Push(0)
	test.Push(3)
	test.Push(0)
	fmt.Println(test.GetMin())
	test.Pop()
	fmt.Println(test.GetMin())
	test.Pop()
	fmt.Println(test.GetMin())
	test.Pop()
	fmt.Println(test.GetMin())
}
