// @Description
// @Author 小游
// @Date 2021/04/07
package q42

func maxSubArray(nums []int) int {
	// 如果长度为0，那么我们直接返回0
	if len(nums) == 0 {
		return 0
	}
	// 设当前值为最大
	max := nums[0]
	// 然后我们开始遍历整个数组
	for i := 1; i < len(nums); i++ {
		// 当前的最大值有两种情况，要么当前是最大，要么是前一个加上当前这个最大
		nums[i] = Max(nums[i-1]+nums[i], nums[i])
		// 然后我们找出最大值就可以了
		if max < nums[i] {
			max = nums[i]
		}
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
