// @Description
// @Author 小游
// @Date 2021/03/31
package q14_1

import "math"

// 解法1 数学推导
func cuttingRope1(n int) int {
	if n <= 3 {
		return n - 1
	}
	a := n / 3
	b := n % 3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	}
	if b == 1 {
		return int(math.Pow(3, float64(a-1))) * 4
	}
	return int(math.Pow(3, float64(a))) * 2
}

// 解法2 使用递归
func cuttingRope2(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	case 5:
		return 6
	case 6:
		return 9
	default:
		return cuttingRope2(n-3) * 3 % 1000000007
	}
}

// 解法3 使用动态规划
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	// 设置dp的初始条件
	dp[2] = 1
	// 我们从3开始进行遍历计算
	for i := 3; i <= n; i++ {
		// 因为我们需要找出一个最大值，所以我们需要遍历dp数组，找出最大的值
		for j := 2; j < i; j++ {
			// 最大的值可能有下面三种情况
			// 1是自己本身就是最大的
			// 2是j*(i-j) 这两段相乘最大
			// 3是j*dp[i-j] 达到最大（i-j这段继续分的最大值）
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
