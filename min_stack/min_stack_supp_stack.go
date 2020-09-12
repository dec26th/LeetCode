package min_stack

import "math"

type MinStackTwo struct {
	stack []int
	MinStackTwo []int
}

/** initialize your data structure here. */
func ConstructorTwo() MinStackTwo {
	return MinStackTwo{
		stack: []int{},
		MinStackTwo: []int{math.MaxInt64},
	}
}

func (this *MinStackTwo) Push(x int)  {
	this.stack = append(this.stack, x)
	top := this.MinStackTwo[len(this.MinStackTwo)-1]
	this.MinStackTwo = append(this.MinStackTwo, min(x, top))
}

func (this *MinStackTwo) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.MinStackTwo = this.MinStackTwo[:len(this.MinStackTwo)-1]
}

func (this *MinStackTwo) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStackTwo) GetMin() int {
	return this.MinStackTwo[len(this.MinStackTwo)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

