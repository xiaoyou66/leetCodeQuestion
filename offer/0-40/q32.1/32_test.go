// @Description
// @Author 小游
// @Date 2021/04/06
package q32_1

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	tree := &TreeNode{Val: 3}
	tree.Left = &TreeNode{Val: 9}
	tree.Right = &TreeNode{Val: 20}
	tree.Right.Left = &TreeNode{Val: 15}
	tree.Right.Right = &TreeNode{Val: 7}
	fmt.Println(levelOrder(tree))
	//fmt.Println(cal(35)+cal(37))
}
