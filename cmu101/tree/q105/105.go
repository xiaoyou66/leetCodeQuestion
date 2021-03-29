// @Description
// @Author 小游
// @Date 2021/03/27
package q105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	hash := make(map[int]int)
	for i := 0; i < len(preorder); i++ {
		hash[inorder[i]] = i
	}
	return buildHelp(hash, preorder, 0, len(preorder)-1, 0)
}

func buildHelp(hash map[int]int, preorder []int, s0 int, e0 int, s1 int) *TreeNode {
	if s0 > e0 {
		return nil
	}
	mid := preorder[s1]
	index := hash[mid]
	leftLen := index - s0 - 1
	node := TreeNode{Val: mid}
	node.Left = buildHelp(hash, preorder, s0, index-1, s1+1)
	node.Right = buildHelp(hash, preorder, index+1, e0, s1+2+leftLen)
	return &node
}
