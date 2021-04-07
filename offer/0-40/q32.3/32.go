// @Description
// @Author 小游
// @Date 2021/04/06
package q32_3

import (
	"container/list"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解法一 层次遍历加双端列队
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 当前层
	queue := list.New()
	// 最终结果
	var res [][]int
	queue.PushBack(root)
	// 用n来表示遍历的方向
	for queue.Len() != 0 {
		var data []int
		// 这里我们把当前队列清空，把当前这一层的数据全部打印出来
		for i := queue.Len(); i > 0; i-- {
			// 同时我们获取队列头部元素
			node := queue.Remove(queue.Front()).(*TreeNode)
			// 判断是奇数层还是偶数层
			if len(res)%2 == 0 {
				// 如res是偶数，说明当前层是奇数层，此时我们要按照从左到右的顺序进行打印
				data = append(data, node.Val)
			} else {
				// 否则我们就按照从右到左的顺序进行打印
				data = append([]int{node.Val}, data...)
			}
			// 添加元素
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

// 解法二 层次遍历加双端列队（我们进行逻辑拆分，避免冗余）
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 当前层
	queue := list.New()
	// 最终结果
	var res [][]int
	queue.PushBack(root)
	// 用n来表示遍历的方向
	for queue.Len() != 0 {
		// 遍历奇数层
		var data []int
		// 这里我们把当前队列清空，把当前这一层的数据全部打印出来
		for i := queue.Len(); i > 0; i-- {
			// 同时我们获取队列头部元素
			node := queue.Remove(queue.Front()).(*TreeNode)
			// 判断是奇数层还是偶数层
			data = append(data, node.Val)
			// 添加元素
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, data)
		// 遍历偶数层
		if queue.Len() == 0 {
			break
		}
		data = []int{}
		// 这里我们把当前队列清空，把当前这一层的数据全部打印出来
		for i := queue.Len(); i > 0; i-- {
			// 同时我们获取队列头部元素
			node := queue.Remove(queue.Back()).(*TreeNode)
			// 判断是奇数层还是偶数层
			data = append(data, node.Val)
			// 添加元素
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
		}
		res = append(res, data)
	}
	return res
}

// 解法三 层次遍历加反转
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 当前层
	queue := list.New()
	// 最终结果
	var res [][]int
	queue.PushBack(root)
	// 用n来表示遍历的方向
	for queue.Len() != 0 {
		var data []int
		// 这里我们把当前队列清空，把当前这一层的数据全部打印出来
		for i := queue.Len(); i > 0; i-- {
			// 同时我们获取队列头部元素
			node := queue.Remove(queue.Front()).(*TreeNode)
			// 判断是奇数层还是偶数层
			// 如res是偶数，说明当前层是奇数层，此时我们要按照从左到右的顺序进行打印
			data = append(data, node.Val)
			// 添加元素
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		// 如果是奇数层，我们就需要进行反转
		if len(res)%2 == 1 {
			sort.Sort(sort.Reverse(sort.IntSlice(data)))
		}
		res = append(res, data)
	}
	return res
}
