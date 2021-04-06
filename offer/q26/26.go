// @Description
// @Author 小游
// @Date 2021/04/05
package q26

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 这里我们先判断两个节点是否都不为空
	// 如果不为空的话，我们尝试判断AB这两颗树是否一样
	// 如果不一样我们再判断A的左子树或者A的右子树和B是否一样
	return (A != nil && B != nil) && (isSame(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

// 判断两颗树是否相等
func isSame(A *TreeNode, B *TreeNode) bool {
	// 当B为空时，此时我们就返回true(因为已经比较完毕了)
	if B == nil {
		return true
	}
	// 如果A节点为空，B节点不为空
	// 或者A B两个节点的值不相等，我们就返回false
	if A == nil || A.Val != B.Val {
		return false
	}
	// 然后我们分别判断左右子树，确保左右子树也符合条件
	return isSubStructure(A.Left, B.Left) && isSubStructure(A.Right, B.Right)
}
