// @Description
// @Author 小游
// @Date 2021/03/27
package q637

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	node := TreeNode{Val: 3}
	node.Left = &TreeNode{Val: 9}
	node.Right = &TreeNode{Val: 20}
	node.Right.Left = &TreeNode{Val: 15}
	node.Right.Right = &TreeNode{Val: 7}
	res := averageOfLevels(&node)
	fmt.Println(res)
}
