// @Description
// @Author 小游
// @Date 2021/04/05
package q22

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next = &ListNode{Val: 4}
	node.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Println(getKthFromEnd(node, 2))
	//fmt.Println(cal(35)+cal(37))
}
