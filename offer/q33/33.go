// @Description
// @Author 小游
// @Date 2021/04/06
package q33

import (
	"container/list"
	"math"
)

// 解法一 使用二叉搜索树
func verifyPostorder1(postorder []int) bool {
	// 这里开始进行判断，我们需要传入数组，然后还有起始点和终止点
	return recur(postorder, 0, len(postorder)-1)
}

// 辅助函数
func recur(postorder []int, i int, j int) bool {
	// 当i大于或者等于j时，我们就比较完毕了，可以返回true了
	if i >= j {
		return true
	}
	// 然后我们记录一下起始点的位置
	p := i
	// j起始就是根节点，我们的左子树比根节点要小，所以我们这里找出符合条件的所有点
	for postorder[p] < postorder[j] {
		p++
	}
	// m暂存一下，此时m恰好为右子树的第一个节点
	m := p
	// 这里我们开始下一轮的比较，这里我们找出所有的右子树
	for postorder[p] > postorder[j] {
		p++
	}
	// 当右子树遍历完时，此时p一定会是根节点（如果不是的话，那就不是搜索二叉树）
	// 此时我们还需要分别比较左子树和右子树是否符合条件
	return p == j && recur(postorder, i, m-1) && recur(postorder, m, j-1)
}

// 解法二 辅助单调栈
func verifyPostorder(postorder []int) bool {
	stack := list.New()
	root := math.MaxInt32
	// 我们后序遍历的方向的节点是 根、右、左 的顺序
	// 所以一开始越往右越大的
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		// 一旦发现我们的节点小于小于栈顶元素时，就说明我们要进入左子树了
		// 接下来，数组继续往前遍历，之后的左子树的每个节点，都要比子树的根要小，才能满足二叉搜索树
		for stack.Len() != 0 && stack.Back().Value.(int) > postorder[i] {
			root = stack.Remove(stack.Back()).(int)
		}
		// 先把值推入栈中
		stack.PushBack(postorder[i])
	}
	return true
}
