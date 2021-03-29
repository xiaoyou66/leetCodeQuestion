// @Description
// @Author 小游
// @Date 2021/03/27
package q21

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	node := ListNode{1, nil}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 4}

	node2 := ListNode{1, nil}
	node2.Next = &ListNode{Val: 3}
	node2.Next.Next = &ListNode{Val: 4}
	res := mergeTwoLists(&node, &node2)
	fmt.Println(res)
}
