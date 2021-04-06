// @Description
// @Author 小游
// @Date 2021/04/05
package q25

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 4}
	node1 := &ListNode{Val: 1}
	node1.Next = &ListNode{Val: 3}
	node1.Next.Next = &ListNode{Val: 4}
	res := mergeTwoLists(node, node1)
	fmt.Println(res)
	//fmt.Println(cal(35)+cal(37))
}
