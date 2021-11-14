package min_stack

import "math"

type MinStack struct {
	container []int
	length    int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		container: *new([]int),
		length:    0,
	}
}

func (this *MinStack) Push(x int) {
	this.container = append(this.container, x)
	this.length++
}

func (this *MinStack) Pop() {
	this.length--
	this.container = this.container[0:this.length]
}

func (this *MinStack) Top() int {
	return this.container[this.length-1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt32
	for _, v := range this.container {
		if min > v {
			min = v
		}
	}

	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
