// Package nc161
// @Description
// @Author 小游
// @Date 2021/04/19
package nc161

import (
	"fmt"
	"testing"
)

func test() {

}

func Test_Tree(t *testing.T) {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(threeOrders(root))
}
