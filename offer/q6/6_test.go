// @Description
// @Author 小游
// @Date 2021/03/29
package q6

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	node := ListNode{Val: 1}
	node.Next = &ListNode{Val: 3}
	node.Next.Next = &ListNode{Val: 2}
	fmt.Println(reversePrint(&node))
}
