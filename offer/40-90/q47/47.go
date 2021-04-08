// @Description
// @Author 小游
// @Date 2021/04/08
package q47

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	// 获取grid的长度和宽度
	m, n := len(grid), len(grid[0])
	// 顶部和左边都是固定的
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	// 然后我们计算其他部分的
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 因为根据我们的状态方程，当前最大值要么是上面的+当前的，要么是左边+当前的，我们选一个最大的即可
			grid[i][j] += max(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
