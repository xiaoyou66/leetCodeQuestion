// @Description
// @Author 小游
// @Date 2021/03/30
package q9

import (
	"container/list"
)

type CQueue struct {
	// stack表示队尾
	stack1 *list.List
	// stack表示队头
	stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	// 如果第二个栈为空，那么我们就把第一个栈的数据全部放进去
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	// 此时我们再判断第二个栈是否为空,不为空那么我们删除第二个栈并返回删除的元素
	if this.stack2.Len() != 0 {
		return this.stack2.Remove(this.stack2.Back()).(int)
	}
	return -1
}
