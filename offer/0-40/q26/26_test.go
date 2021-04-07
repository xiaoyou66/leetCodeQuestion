// @Description
// @Author 小游
// @Date 2021/04/05
package q26

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	node1 := &TreeNode{Val: 3}
	node1.Left = &TreeNode{Val: 4}
	node1.Right = &TreeNode{Val: 5}
	node1.Left.Left = &TreeNode{Val: 1}
	node1.Left.Right = &TreeNode{Val: 2}

	node2 := &TreeNode{Val: 4}
	node2.Left = &TreeNode{Val: 1}
	fmt.Println(isSubStructure(node1, node2))
	//fmt.Println(cal(35)+cal(37))
}
