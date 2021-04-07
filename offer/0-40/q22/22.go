// @Description
// @Author 小游
// @Date 2021/04/05
package q22

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一 两次循环
func getKthFromEnd1(head *ListNode, k int) *ListNode {
	root := head
	// n表示链表长度
	n := 0
	// 第一遍for循环，判断链表长度
	for head != nil {
		n++
		head = head.Next
	}
	// 第二遍，返回倒数倒数第k个节点
	for i := 0; i < n-k; i++ {
		root = root.Next
	}
	// 返回节点
	return root
}

// 解法二 双指针
func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 定义快慢两个节点
	slow, fast := head, head
	t := 0
	for fast != nil {
		// 只有当t>=k的时候我们才移动慢指针，因为我们的t相当于快指针，我们的慢指针要比快指针要小于k
		if t >= k {
			slow = slow.Next
		}
		// 移动快指针
		fast = fast.Next
		t++
	}
	// 判断越界问题
	if t < k {
		return nil
	}
	return slow
}
