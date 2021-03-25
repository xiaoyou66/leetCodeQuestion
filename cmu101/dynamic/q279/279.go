// @Description
// @Author 小游
// @Date 2021/03/19
package q279

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = Min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func Min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
