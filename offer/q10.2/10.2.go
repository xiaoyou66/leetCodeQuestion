// @Description
// @Author 小游
// @Date 2021/03/30
package q10_2

// 动态规划原始版本
func numWays1(n int) int {
	if n <= 1 {
		return 1
	}
	// 为了简单直观，先用dp数组
	dp := make([]int, n+1)
	// 因为我们需要两个初始条件，所以需要先初始化
	dp[1] = 1
	dp[2] = 2
	// 这个题目用动态规划来做，所以我们只需要实现状态转移方程就行了
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % (1e9 + 7)
	}
	return dp[n]
}

// 动态规划压缩空间
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	// 我们使用a代表第0级，b代表第1级,sum表示总可能数
	a := 1
	b := 1
	sum := 0
	// 下面我们就可以计算各级的情况了
	for i := 2; i <= n; i++ {
		// 下面这部分没啥好说的，其实就是当前情况等于n-2+n-1这两个和
		sum = (a + b) % (1e9 + 7)
		a = b
		b = sum

	}
	return sum
}
