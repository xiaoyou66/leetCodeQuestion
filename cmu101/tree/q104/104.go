// @Description
// @Author 小游
// @Date 2021/03/27
package q104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 分别计算左右子树的高度
	l := maxDepth(root.Left) + 1
	r := maxDepth(root.Right) + 1
	if l > r {
		return l
	} else {
		return r
	}
}
