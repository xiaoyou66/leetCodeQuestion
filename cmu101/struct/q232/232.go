// @Description
// @Author 小游
// @Date 2021/03/26
package q232

type Stack struct {
	data []int
}

func (s *Stack) Push(k int) {
	s.data = append(s.data, k)
}
func (s *Stack) Pop() (x int) {
	x = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return
}
func (s *Stack) Empty() bool {
	return len(s.data) == 0
}
func (s *Stack) Top() int {
	return s.data[len(s.data)-1]
}

type MyQueue struct {
	// 我们需要定义一个in，一个out两个栈
	in  Stack
	out Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	// 把当前数据放入in的栈中
	this.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	// 当我们进行pop操作时，我们把in栈里面的数据全部pop到out栈里
	this.in2Out()
	return this.out.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.in2Out()
	// 此时out就是我们队列的结尾了
	return this.out.Top()
}
func (this *MyQueue) in2Out() {
	// 这里是
	if this.out.Empty() {
		for !this.in.Empty() {
			x := this.in.Pop()
			this.out.Push(x)
		}
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.in.Empty() && this.out.Empty()
}
