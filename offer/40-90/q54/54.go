// @Description
// @Author 小游
// @Date 2021/04/09
package q54

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res = 0
var n int

func kthLargest(root *TreeNode, k int) int {
	n = k
	// 开始遍历dfs
	dfs(root)
	return res
}

// dfs深度遍历
func dfs(root *TreeNode) {
	// 如果节点为空，那么就直接返回
	if root == nil {
		return
	}
	// 从右节点开始遍历
	dfs(root.Right)
	// n--操作，这里相当于在统计
	n--
	// 如果n变成0了，说明我们找到第k个最大的数了，这个时候我们直接设置结果
	if n == 0 {
		res = root.Val
	}
	// 然后遍历左子树
	dfs(root.Left)
}
