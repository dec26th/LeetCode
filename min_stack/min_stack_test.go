package min_stack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	test := Constructor()
	test.Push(-2)
	test.Push(0)
	test.Push(-1)
	fmt.Println(test.Min())
	fmt.Println(test.Top())
	test.Pop()
	fmt.Println(test.Min())
}
