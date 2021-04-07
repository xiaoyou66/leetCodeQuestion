// @Description
// @Author 小游
// @Date 2021/04/05
package q25

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一 头节点
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil && l2 != nil {
		// 比较判断
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	// 判断两个节点是否为空
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return res.Next
}

// 方法二 使用递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 如果一个节点为空的情况下，我们返回另外一个非空节点
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
