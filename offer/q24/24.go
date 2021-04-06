// @Description
// @Author 小游
// @Date 2021/04/05
package q24

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法1，使用递归
func reverseList1(head *ListNode) *ListNode {
	// 当我们的head为空的时候我们停止循环,这个时候返回我们的头节点（其实就是返回尾结点）
	if head == nil || head.Next == nil {
		return head
	}
	// 开始递归获取尾结点
	node := reverseList1(head.Next)
	// 此时我们修改一下指向的顺序
	// 即当前指针的下一个指针的下一个指针指向当前指针
	head.Next.Next = head
	// 此时头节点相当于尾结点了，我们需要置为nil
	head.Next = nil
	return node
}

// 解法二 使用双指针
func reverseList(head *ListNode) *ListNode {
	// 表示前一个指针
	var prev *ListNode
	// curr表示当前指针
	curr := head
	for curr != nil {
		// 先记录下一个指针
		next := curr.Next
		// 反转一下位置
		curr.Next = prev
		// 移动一下指针
		prev = curr
		curr = next
	}
	return prev
}
