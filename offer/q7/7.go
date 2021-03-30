// @Description
// @Author 小游
// @Date 2021/03/30
package q7

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解法一 使用递归
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 创建根节点，注意前序遍历的第一个一定是根节点
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	// 从前序序列中去找，找到后序遍历的根节点
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// 前面我们找到了后续遍历的根节点，那么根节点前面的都是左子树，后面的是右子树
	// 这样我们就可以重建这棵二叉树了
	root.Left = buildTree1(preorder[1:i+1], inorder[:i])
	root.Right = buildTree1(preorder[i+1:], inorder[i+1:])
	return root
}

// 解法二 迭代
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 创建一个根节点
	root := &TreeNode{preorder[0], nil, nil}
	// 初始化一个栈
	stack := []*TreeNode{}
	// 先把根节点放入栈中
	stack = append(stack, root)
	var inorderIndex int
	// 遍历一下前序遍历的节点
	for i := 1; i < len(preorder); i++ {
		// 获取前序遍历节点的值
		preorderVal := preorder[i]
		// 然后我们也获取一下栈顶的元素
		node := stack[len(stack)-1]
		// 当我们当前节点的值不等于中序遍历的节点时（此时节点是左子树）
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			// 这里就是当前节点的值等于中序遍历节点的值
			// 此时当前节点就是根节点，但是我们不需要根节点，我们需要右子树
			// 所以这里我们就是在不断出栈，找到不等于根节点的值，此时这个节点就是右节点
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
