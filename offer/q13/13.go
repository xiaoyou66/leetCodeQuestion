// @Description
// @Author 小游
// @Date 2021/03/30
package q13

// 定义全局变量
var n1, m1, k1 int
var visited [][]bool

func movingCount(m int, n int, k int) int {
	n1 = n
	m1 = m
	k1 = k
	// 初始化visited数组
	visited = make([][]bool, m)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	// 进行计算
	return dfs(0, 0, 0, 0)
}

func dfs(i int, j int, si int, sj int) int {
	if i >= n1 || j >= m1 || si+sj > k1 {
		return 0
	}
	// 计算和
	//si1 := cal()
	return 0
}

// 计算数位和
func cal(i int, n int) int {
	if (i+1)%10 == 0 {
		return n - 8
	} else {
		return n + 1
	}
}
