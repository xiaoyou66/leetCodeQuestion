// @Description
// @Author 小游
// @Date 2021/03/19
package q64

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 首先我们可以计算第一行和第一列
	// 因为我们只能向下或者向右移动所以这一行和第一列是前面的累加
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	// 遍历计算后面的dp值
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 第i和j的位置最小值其实就是上面或者左边 中选一个最小的
			grid[i][j] = min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
