// @Description
// @Author 小游
// @Date 2021/04/09
package q55_1

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一 递归
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth1(root.Left), maxDepth1(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 方法二 BFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	count := 0
	// 先把根节点放入到队列中
	queue.PushBack(root)
	// 不断遍历队列直到出队
	for queue.Len() != 0 {
		// 不断遍历让队列出队
		size := queue.Len()
		for size > 0 {
			// 每遍历一次 size就减一
			size--
			// 获取队头元素
			node := queue.Remove(queue.Front()).(*TreeNode)
			// 判断头尾
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}
		count++
	}
	return count
}
