// @Description
// @Author 小游
// @Date 2021/03/31
package q18

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	root := &ListNode{Val: 4}
	root.Next = &ListNode{Val: 5}
	root.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next = &ListNode{Val: 9}
	fmt.Println(deleteNode(root, 5))
	//fmt.Println(cal(35)+cal(37))
}
