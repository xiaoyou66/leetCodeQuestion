// @Description
// @Author 小游
// @Date 2021/04/09
package q55_2

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解法一 后序遍历
func isBalanced1(root *TreeNode) bool {
	// 我们-1表示这棵树不是平衡二叉树
	return recur(root) != -1
}

// 声明一个辅助函数
func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 分别遍历左右子树，获取左右子树的高度差
	left := recur(root.Left)
	// 如果发现左子树的高度差为-1，说明不是平衡二叉树直接返回
	if left == -1 {
		return -1
	}
	// 遍历右子树
	right := recur(root.Right)
	// 如果发现左子树的高度差为-1，说明不是平衡二叉树直接返回
	if right == -1 {
		return -1
	}
	// 然后我们判断一下左右子树的高度差
	if math.Abs(float64(left-right)) < 2 {
		// 小于2就说明在范围内，此时我们返回最深的数的高度
		return max(left, right) + 1
	} else {
		return -1
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 解法二 先序遍历
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 首先我们要确保当前树的左右子树深度不超过2
	// 并且我们还要确保左右子树也是平衡二叉树
	return math.Abs(float64(depth(root.Left)-depth(root.Right))) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

// 判断二叉树的深度
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归左右子树，返回最大值
	left := depth(root.Left)
	right := depth(root.Right)
	// 返回最大深度
	return max(left, right) + 1
}
