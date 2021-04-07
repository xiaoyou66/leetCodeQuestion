// @Description
// @Author 小游
// @Date 2021/03/30
package q13

import "container/list"

// 定义全局变量
var n1, m1, k1 int
var visited [][]bool

func movingCount1(m int, n int, k int) int {
	m1 = m
	n1 = n
	k1 = k
	// 初始化访问数组
	visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	// dfs遍历
	return dfs(0, 0, 0, 0)
}

// si和sj分别表示位数之和
func dfs(i int, j int, si int, sj int) int {
	if i >= m1 || j >= n1 || si+sj > k1 || visited[i][j] {
		return 0
	}
	// 标记为访问状态
	visited[i][j] = true
	// sj1表示往右走
	sj1 := cal(j, sj)
	// sj1表示往下走
	si1 := cal(i, si)
	// 然后把所有的情况都加起来就是机器人能到达的范围了
	return 1 + dfs(i, j+1, si, sj1) + dfs(i+1, j, si1, sj)
}

// 计算数位和
func cal(i int, n int) int {
	if (i+1)%10 == 0 {
		return n - 8
	} else {
		return n + 1
	}
}

// 解法二 广度优先遍历
func movingCount(m int, n int, k int) int {
	// 初始化一个队列
	queue := list.List{}
	// 初始化访问数组
	visited := make([][]bool, m)
	// 数据开辟空间
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	// 队列里面放入四个初始值（分别代表i j si sj这四个位置）
	queue.PushBack([]int{0, 0, 0, 0})
	res := 0
	// 当我们队列长度等于0时退出循环
	for queue.Len() > 0 {
		// 获取队列顶的元素
		bfs := queue.Remove(queue.Back()).([]int)
		// 赋值
		i := bfs[0]
		j := bfs[1]
		si := bfs[2]
		sj := bfs[3]
		// 先判断当前值是否符合条件
		if i >= m || j >= n || si+sj > k || visited[i][j] {
			continue
		}
		// 步数+1，设置当前位置访问
		res++
		visited[i][j] = true
		// 计算右下两个位置
		sj1 := cal(j, sj)
		si1 := cal(i, si)
		// 把我们遍历的点放入栈中
		queue.PushBack([]int{i + 1, j, si1, sj})
		queue.PushBack([]int{i, j + 1, si, sj1})
	}
	return res
}
