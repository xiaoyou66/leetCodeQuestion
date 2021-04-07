// @Description
// @Author 小游
// @Date 2021/03/31
package q14_2

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n] % (1e9 + 7)
}

func max(i, j int) int {
	if i > j {
		return i % (1e9 + 7)
	}
	return j % (1e9 + 7)
}
