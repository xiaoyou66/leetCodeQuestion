// Package nc161
// @Description
// @Author 小游
// @Date 2021/04/19
package nc161

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 开始遍历
func threeOrders(root *TreeNode) [][]int {
	ans := [][]int{{}, {}, {}}
	PreMidLastRoot(root, &ans[0], &ans[1], &ans[2])
	return ans
}

// PreMidLastRootStack 不用递归，用栈模拟递归
func PreMidLastRootStack(root *TreeNode) (pre []int, mid []int, last []int) {
	var stk1 []*TreeNode
	var stk2 []int //stk1[i]对应的stk2[i], stk2[i]为 1：表示放入，前序； 2表示取出判断左子树， 3表示取出判断右子树
	if root == nil {
		return nil, nil, nil
	}
	// 放入栈时，前序pre里加入Val
	// 第一次拿出看左子树时，pre加入左子树
	// 第二次拿出看右子树，mid加入。同时pre加入右子树
	// 最后一次拿出时， last加入
	stk1 = append(stk1, root)
	stk2 = append(stk2, 1)
	pre = append(pre, root.Val)
	for len(stk1) > 0 {
		top := len(stk1) - 1
		cur := stk1[top]
		if stk2[top] == 1 {
			stk2[top] = 2
			if cur.Left != nil {
				pre = append(pre, cur.Left.Val)
				stk1 = append(stk1, cur.Left)
				stk2 = append(stk2, 1)
			}
		} else if stk2[top] == 2 {
			stk2[top] = 3
			mid = append(mid, cur.Val)
			if cur.Right != nil {
				pre = append(pre, cur.Right.Val)
				stk1 = append(stk1, cur.Right)
				stk2 = append(stk2, 1)
			}
		} else {
			last = append(last, cur.Val)
			stk1 = stk1[:top]
			stk2 = stk2[:top]
		}
	}
	return
}

// PreMidLastRoot 一次递归
func PreMidLastRoot(root *TreeNode, pre, mid, last *[]int) {
	if root == nil {
		return
	}
	*pre = append(*pre, root.Val)
	PreMidLastRoot(root.Left, pre, mid, last)
	*mid = append(*mid, root.Val)
	PreMidLastRoot(root.Right, pre, mid, last)
	*last = append(*last, root.Val)
}

// 先序遍历
func preNode(root *TreeNode, data *[]int) []int {
	if root == nil {
		return *data
	}
	*data = append(*data, root.Val)
	preNode(root.Left, data)
	preNode(root.Right, data)
	return *data
}

// 中序遍历
func midNode(root *TreeNode, data *[]int) []int {
	if root == nil {
		return *data
	}
	midNode(root.Left, data)
	*data = append(*data, root.Val)
	midNode(root.Right, data)
	return *data
}

// 后序遍历
func lastNode(root *TreeNode, data *[]int) []int {
	if root == nil {
		return *data
	}
	lastNode(root.Left, data)
	lastNode(root.Right, data)
	*data = append(*data, root.Val)
	return *data
}

/* 不用递归，用栈模拟递归,使用1个栈。精简版*/
type Node struct {
	Root *TreeNode
	Mark int
}

func PreMidLastRootStack2(root *TreeNode) (pre []int, mid []int, last []int) {
	var stk []Node
	if root == nil {
		return nil, nil, nil
	}
	//初始把root放入栈，makr = 1. pre,mid, last都为Nil
	//第一次拿出时makr=1,val放入pre,同时看左子树不为Nil则加入栈。
	//第二次拿出时mark=2,val放入mid,同时看右子树不为Nil则加入栈。
	//第三次拿出时mark=3,val放入last,同时pop栈顶。因为左右子树都遍历过了。
	//每次放入栈时，mark都是1.出栈时mark=3
	stk = append(stk, Node{
		Root: root,
		Mark: 1,
	})
	for len(stk) > 0 {
		top := len(stk) - 1
		node := &stk[top]
		if node.Mark == 1 {
			node.Mark = 2
			pre = append(pre, node.Root.Val)
			if node.Root.Left != nil {
				stk = append(stk, Node{
					Root: node.Root.Left,
					Mark: 1,
				})
			}
		} else if node.Mark == 2 {
			node.Mark = 3
			mid = append(mid, node.Root.Val)
			if node.Root.Right != nil {
				stk = append(stk, Node{
					Root: node.Root.Right,
					Mark: 1,
				})
			}
		} else {
			last = append(last, node.Root.Val)
			stk = stk[:top]
		}
	}
	return
}
