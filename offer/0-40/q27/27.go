// @Description
// @Author 小游
// @Date 2021/04/05
package q27

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree1(root *TreeNode) *TreeNode {
	// 如果节点为空，那么我们就返回空值
	if root == nil {
		return root
	}
	// 遍历左右节点
	right := mirrorTree1(root.Right)
	left := mirrorTree1(root.Left)
	// 交换在左右节点
	root.Left = right
	root.Right = left
	return root
}

func mirrorTree(root *TreeNode) *TreeNode {
	// 如果根节点为空，那么我们就返回空
	if root == nil {
		return nil
	}
	// 初始化一个栈
	stack := list.New()
	// 元素入栈操作
	stack.PushBack(root)
	// 当我们栈为空时表示我们替换结束
	for stack.Len() != 0 {
		// 元素出栈
		node := stack.Remove(stack.Back()).(*TreeNode)
		// 然后我们判断左右子树是否为空
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		// 交换一下元素
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
	}
	return root
}
