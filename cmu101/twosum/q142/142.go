// @Description 
// @Author 小游
// @Date 2021/03/12
package q142

type ListNode struct {
	Val int
    Next *ListNode
}

// go里面的do while循环参考 https://yourbasic.org/golang/do-while-loop/
// 链表找环这个问题用快慢指针是通用解法，为什么两个指针一定会相遇呢？ 其实快慢指针一个看相对静止，另一个是运动的，所以这两个追的上 参考https://www.zhihu.com/question/23208893
// 然后为什么fast到头节点两个移动为什么相遇的点一定是环呢？ 
func detectCycle(head *ListNode) *ListNode {
	slow,fast:=head,head
	for  {
		// 首先必须要判断fast和fast.next指针不能为空
		if fast == nil || fast.Next == nil {
			return nil
		}
		// fast前进两次，slow前进一次
		fast = fast.Next.Next
		slow = slow.Next
		// 当我们的fast指针和slow指针相遇时，我们跳出循环
		if slow == fast {
			break
		}
	}
	// 此时，我们的fast指针回到原点
	fast = head
	// 然后我们继续遍历直到fast和slow再次相遇,这个点就是我们要找的点
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}