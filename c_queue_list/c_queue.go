package c_queue_list

import "container/list"

type CQueue struct {
	stack1 *list.List
}

func Constructor () CQueue {
	return CQueue{
		stack1: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack1.Len() == 0 {
		 return -1
	}

	e := this.stack1.Front()
	this.stack1.Remove(e)

	return e.Value.(int)
}
