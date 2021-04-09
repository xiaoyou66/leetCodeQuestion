// @Description
// @Author 小游
// @Date 2021/04/09
package q52

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	dic := make(map[*ListNode]bool)
	// 首先我们遍历第一个节点，把节点所有的值全部放入map中
	for headA != nil {
		dic[headA] = true
		headA = headA.Next
	}
	// 第二次遍历headB
	for headB != nil {
		if _, ok := dic[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}
