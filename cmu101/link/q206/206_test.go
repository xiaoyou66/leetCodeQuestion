// @Description
// @Author 小游
// @Date 2021/03/27
package q206

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	node := ListNode{1, nil}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next = &ListNode{Val: 4}
	node.Next.Next.Next.Next = &ListNode{Val: 5}
	res := reverseList(&node)
	fmt.Println(res)
}
