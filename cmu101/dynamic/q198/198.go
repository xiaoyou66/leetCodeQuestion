// @Description
// @Author 小游
// @Date 2021/03/17
package q198

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	a, b, c := 0, 0, 0
	n := len(nums)
	// 这里我们直接从0开始计算
	for i := 0; i < n; i++ {
		if b > nums[i]+c {
			a = b
		} else {
			a = c + nums[i]
		}
		c = b
		b = a
	}
	return a
}
