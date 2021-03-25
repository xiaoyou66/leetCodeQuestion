// @Description
// @Author 小游
// @Date 2021/03/19
package q542

import "math"

func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	// 初始化数组
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 为了方便，我们先把所有的都设为int的最大值，方便后面比较
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	// 这个东西的状态转移方程比较简单就是从上下左右四个反向中找到一个最小的
	// 不过我们第一次遍历是从左上扫描到右下，所以这里我们只能获取上和左这两个方向
	// 所以我们这里需要扫描两次，完成上下左右这四个方向
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果当前位置为0，那么结果就是0
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				// 在确定最小值时我们先确保上和左可以访问并和自身比较
				// 这里为什么要和自身比较呢，因为一开始我们dp值就是最大的
				// 然后这里我们是分别比较了上和左，所以我们需要使用dp[i][j]来临时存储最小的值
				// 最后就是我们第二次扫描时，可能第一次扫描时当前值已经是最小的了，所以我们就没有必要非要从另外两个方向比较
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
				}
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
				}
			}
		}
	}
	// 下面这个就是从右下扫描到左上，找出下和右中最小的值
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 因为我们第一次扫描的时候已经知道0了，所以这里没有必要再次比较
			if matrix[i][j] != 0 {
				if i+1 < m {
					dp[i][j] = min(dp[i+1][j]+1, dp[i][j])
				}
				if j+1 < n {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
			}
		}
	}
	return dp
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
