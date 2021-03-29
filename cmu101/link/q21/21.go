// @Description
// @Author 小游
// @Date 2021/03/27
package q21

type ListNode struct {
	Val  int
	Next *ListNode
}

// 非递归方法
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	node := dummy
	// 不断进行遍历,有一个为空，我们就退出循环
	for l1 != nil && l2 != nil {
		// 找到一个大的值进行拼接
		if l1.Val >= l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		node = node.Next
	}
	// 把未接上的节点接到当前节点的后面去
	if l1 != nil {
		node.Next = l1
	} else if l2 != nil {
		node.Next = l2
	}
	return dummy.Next
}

// 递归解法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 有一个为空时，我们直接返回另一个就行了
	if l2 == nil {
		return l1
	} else if l1 == nil {
		return l2
	}
	// 当l1当前的值大于l2时，我们接上l2的剩余部分
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
	// 把l1的剩余部分合并一下
	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}
