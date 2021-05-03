// Package important @Description
// @Author 小游
// @Date 2021/04/11
package important

import (
	"math"
)

// 最长公共子串
func getLCS(a string, b string) int {
	m, n := len(a), len(b)
	// 创建dp数组
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// 最大值max
	max := 0
	// 计算dp数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果两个字符串相等，那么当前dp数组为上一个+1
			if a[i] == b[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				max = Max(max, dp[i+1][j+1])
			}
		}
	}
	return max
}

func getLCS2(a string, b string) int {
	m, n := len(a), len(b)
	// 创建dp数组
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// 计算dp数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果两个字符串相等，那么当前dp数组为上一个+1
			if a[i] == b[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = Max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 凑硬币问题
// 凑硬币问题
func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	data := make([]int, amount)
	return aux(coins, amount, data)
}

// 定义一个辅助函数来凑零钱
// 我们传入硬币数，需要凑齐金钱数，以及我们定义的一个记忆函数
func aux(coins []int, rem int, count []int) int {
	if rem < 0 {
		return -1
	} else if rem == 0 {
		return 0
	} else if count[rem-1] != 0 {
		return count[rem-1]
	}
	min := math.MaxInt32
	for _, v := range coins {
		tmp := aux(coins, rem-v, count)
		if tmp >= 0 && tmp < min {
			min = tmp + 1
		}
	}
	if min == math.MaxInt32 {
		count[rem-1] = -1
	} else {
		count[rem-1] = min
	}
	return count[rem-1]
}

// 动态规划
func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for k := range dp {
		dp[k] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				if dp[i-coins[j]]+1 < dp[i] {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
