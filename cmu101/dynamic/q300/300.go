// @Description https://leetcode-cn.com/problems/longest-increasing-subsequence/
// @Author 小游
// @Date 2021/03/q20
package q300

func lengthOfLIS(nums []int) int {
	max, n := 0, len(nums)
	if n <= 1 {
		return n
	}
	// 初始情况，我们设置所有dp的数组为1
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	// 这里我们需要使用两层循环来进行遍历
	for i := 0; i < n; i++ {
		// 这里一层for循环是以为我们这个是求递增的子序列
		// 为了获取到所有情况，所以我们这里使用dp数组来实现
		for j := 0; j < i; j++ {
			// 只要i的值大于j，这个时候我们就可以说序列是递增的
			if nums[i] > nums[j] {
				// 然后我们只需要比较一下，找出最大的一个就可以了
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		// 这里我们实时更新一下max的值
		max = Max(max, dp[i])
	}
	return max
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
