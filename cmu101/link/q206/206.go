// @Description
// @Author 小游
// @Date 2021/03/27
package q206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 下面这个是递归写法
	// return reverse(head,nil)
	var pre *ListNode
	var next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre

		pre = head
		head = next
	}
	return pre
}

// 递归法翻转链表，这里我们需要传入头节点和前一个节点
func reverse(head *ListNode, pre *ListNode) *ListNode {
	// 当头节点为空时，我们返回前一个节点
	if head == nil {
		return pre
	}
	// 先记录下一个节点
	next := head.Next
	// 这里我们知道前一个节点了，下面我们就翻转一下
	head.Next = pre
	// 最后我们返回的节点
	return reverse(next, head)
}
