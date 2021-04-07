// @Description
// @Author 小游
// @Date 2021/03/29
package q6

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint2(head *ListNode) []int {
	data := new([]int)
	reverse(data, head)
	return *data
}

// 使用辅助函数进行递归
func reverse(data *[]int, head *ListNode) {
	// 当头部为空时直接返回即可
	if head == nil {
		return
	}
	// 这里就是关键了，因为我们是需要反过来遍历的
	// 所以其实我们这里就是先递归到链表终点
	reverse(data, head.Next)
	// 然后我们再把当前值放入数组中
	*data = append(*data, head.Val)
}

// 解法二 使用栈
func reversePrint(head *ListNode) []int {
	// 新建一个栈
	stack := list.New()
	// 遍历这个链表，把数据压入栈中
	for head != nil {
		stack.PushFront(head.Val)
		head = head.Next
	}
	// 遍历栈，把所有的数据弹出，并放入数组
	var data []int
	for e := stack.Front(); e != nil; e = e.Next() {
		data = append(data, e.Value.(int))
	}
	// 返回数据
	return data
}
