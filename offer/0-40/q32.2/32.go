// @Description
// @Author 小游
// @Date 2021/04/06
package q32_2

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 我的解法，使用两个队列
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 当前层
	queue := list.New()
	// 最终结果
	var res [][]int
	queue.PushBack(root)
	for queue.Len() != 0 {
		var data []int
		// 初始化一个新的队列来暂存
		queue2 := list.New()
		// 这里我们把当前队列清空，把当前这一层的数据全部打印出来
		for queue.Len() != 0 {
			// 同时我们获取队列头部元素
			node := queue.Remove(queue.Front()).(*TreeNode)
			data = append(data, node.Val)
			if node.Left != nil {
				queue2.PushBack(node.Left)
			}
			if node.Right != nil {
				queue2.PushBack(node.Right)
			}
		}
		res = append(res, data)
		// 切换队列
		queue = queue2
	}
	return res
}

// 大佬的解法，使用一个队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 当前层
	queue := list.New()
	// 最终结果
	var res [][]int
	queue.PushBack(root)
	for queue.Len() != 0 {
		fmt.Println(queue.Len())
		var data []int
		// 这里我们仅临时记录一下当前队列的长度，然后把元素全部放入数组中
		tmp := queue.Len()
		for i := 0; i < tmp; i++ {
			// 同时我们获取队列头部元素
			node := queue.Remove(queue.Front()).(*TreeNode)
			data = append(data, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, data)
	}
	return res
}
