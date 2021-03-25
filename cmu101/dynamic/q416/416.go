// @Description
// @Author 小游
// @Date 2021/03/20
package q416

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, max := 0, 0
	// 这里我们计算总数的时候还需要顺便计算最大值
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	// 判断是否可以平分为两部分
	if sum%2 != 0 {
		return false
	}
	// 计算我们的目标值
	target := sum / 2
	// 当最大值大于目标值时，一定是不能再分的
	if max > target {
		return false
	}
	// 初始化一个dp数组，一维对应数组的个数，二维对应我们背包问题的目标值
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	// 默认情况下，当我们的目标值为0时，一定可以满足
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	// 然后就是当i为0时，我们只能存放nums[0]这个元素，此时当我们的目标值为num[0]时，也满足条件
	dp[0][nums[0]] = true
	// 这里我们开始计算dp数组（注意0已经计算好了，我们从1开始）
	for i := 1; i < n; i++ {
		// 获取当前的值
		v := nums[i]
		// 从1开始挨个计算目标值为j的情况下，放i个元素可能的结果
		for j := 1; j <= target; j++ {
			// 当我们目标值大于i当前值时
			if j >= v {
				// 我们可以判断一下如果i小于1，是否可以满足条件
				// 或者前面一个恰好放入v时满足条件
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}
