// @Description
// @Author 小游
// @Date 2021/04/05
package q27

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	node1 := &TreeNode{Val: 4}
	node1.Left = &TreeNode{Val: 2}
	node1.Right = &TreeNode{Val: 7}
	node1.Left.Left = &TreeNode{Val: 1}
	node1.Left.Right = &TreeNode{Val: 3}
	node1.Right.Left = &TreeNode{Val: 6}
	node1.Right.Right = &TreeNode{Val: 9}
	res := mirrorTree(node1)
	fmt.Println(res)
	//fmt.Println(cal(35)+cal(37))
}
