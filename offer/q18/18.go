// @Description
// @Author 小游
// @Date 2021/03/31
package q18

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	root := head
	pre := head
	// 找到要删除节点的前一个节点
	for head != nil && head.Val != val {
		pre = head
		head = head.Next
	}
	// 判断是否为根节点
	if root.Val == val {
		return root.Next
	}
	// 如果不是那么就把节点删除即可
	if head != nil {
		pre.Next = pre.Next.Next
	}
	return root
}
