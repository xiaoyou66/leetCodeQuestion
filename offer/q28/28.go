// @Description
// @Author 小游
// @Date 2021/04/05
package q28

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 判断左右两边是否对称
	return recur(root.Right, root.Left)
}

func recur(l *TreeNode, r *TreeNode) bool {
	// 当左右两边都为空的时候我们返回 true
	if l == nil && r == nil {
		return true
	}
	// 当两个节点不相同或者有一个节点为空时，我们退出循环
	if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	// 这里我们简单的判断两端是否对称
	return recur(l.Left, r.Right) && recur(l.Right, r.Left)
}
