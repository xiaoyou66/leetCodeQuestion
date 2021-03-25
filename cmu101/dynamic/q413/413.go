// @Description
// @Author 小游
// @Date 2021/03/19
package q413

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	// 初始化dp数组
	dp := make([]int, len(nums))
	// 遍历我们的nums，来计算dp数组
	for i := 2; i < len(nums); i++ {
		// 当前位置是否符合等差数列的定义
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			// 如果满足，那么dp就为上个dp+1
			dp[i] = dp[i-1] + 1
		}
	}
	sum := 0
	// 计算dp数组和
	for _, v := range dp {
		sum += v
	}
	return sum
}
