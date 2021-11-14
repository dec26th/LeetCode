package c_queue

// 剑指 Offer 09 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
// 分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

type CQueue struct {
	originStack []int
}

func Constructor() CQueue {
	return CQueue{
		originStack: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.originStack = append(this.originStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.originStack) == 0 {
		return -1
	}
	supStack := make([]int, len(this.originStack)-1)

	for i := len(this.originStack) - 1; i > 0; i-- {
		supStack[len(this.originStack)-i-1] = this.originStack[i]
	}
	result := this.originStack[0]

	this.originStack = make([]int, len(supStack))

	for i := 0; i < len(this.originStack); i++ {
		this.originStack[i] = supStack[len(this.originStack)-1-i]
	}
	return result
}
