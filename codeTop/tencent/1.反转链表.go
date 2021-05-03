// Package tencent
// @Description
// @Author 小游
// @Date 2021/04/15
package tencent

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一递归
func reverseList1(head *ListNode) *ListNode {
	// 先判断节点是否为空(如果节点的后一位为空，说明我们已经到达节点末尾了)
	if head == nil || head.Next == nil {
		return head
	}
	// 这里是理解的关键，这个递归他的目的就是直接到达链表的末尾
	res := reverseList1(head.Next)
	// 这里我们可以理解为head已经是尾结点了，所以我们需要修改一下方向
	head.Next.Next = head
	head.Next = nil
	return res
}

// 解法二迭代
func reverseList(head *ListNode) *ListNode {
	// 设置前一个节点
	var prev *ListNode = nil
	// 然后设置当前节点
	var curr = head
	// 当当前的节点为空的时候我们跳出循环
	for curr != nil {
		// 我们先暂时保存下一个节点
		tmp := curr.Next
		// 当前节点的下一个节点就是前一个节点
		curr.Next = prev
		// 修改后我们就把前一个接地那切换为当前节点
		prev = curr
		// 然后当前节点就是下一个基点
		curr = tmp
	}
	return prev
}
