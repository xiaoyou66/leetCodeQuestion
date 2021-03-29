// @Description
// @Author 小游
// @Date 2021/03/27
package q637

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Queue struct {
	data []*TreeNode
}

func (q *Queue) Push(k *TreeNode) {
	q.data = append(q.data, k)
}
func (q *Queue) Pop() (i *TreeNode) {
	i = q.data[0]
	q.data = q.data[1:]
	return
}
func (q *Queue) Front() (i *TreeNode) {
	return q.data[0]
}
func (q *Queue) Empty() bool {
	return len(q.data) == 0
}
func averageOfLevels(root *TreeNode) []float64 {
	var ans []float64
	// 节点为空时，我们直接返回结果
	if root == nil {
		return ans
	}
	// 初始化队列
	q := Queue{}
	// 先把当前节点入队
	q.Push(root)
	// 当q为空时退出
	for !q.Empty() {
		// 当前队列的大小
		count := len(q.data)
		// 队列和
		sum := 0
		for i := 0; i < count; i++ {
			// q出队
			node := q.Pop()
			// 然后计算和
			sum += node.Val
			// 下面我们就把下一层的节点推入队列(包括左右节点)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		// 计算当前层的平均值
		ans = append(ans, float64(sum)/float64(count))
	}
	return ans
}
