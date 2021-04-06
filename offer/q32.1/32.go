// @Description
// @Author 小游
// @Date 2021/04/06
package q32_1

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// 我们使用队列来进行操作
	queue := list.New()
	queue.PushBack(root)
	// 初始化结果数组
	var res []int
	// 当栈为空时，我们停止循环
	for queue.Len() != 0 {
		// 打印当前值并出栈
		node := queue.Remove(queue.Front()).(*TreeNode)
		res = append(res, node.Val)
		// 判断左右子树是否为空,如果不为空，我们就进栈
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return res
}
