// @Description
// @Author 小游
// @Date 2021/03/25
package q448

import "math"

func findDisappearedNumbers(nums []int) []int {
	var ans []int
	// 遍历nums数组
	for _, v := range nums {
		// 首先我们计算当前数字所在的位置（数组从0开始，所以我们需要-1）
		pos := int(math.Abs(float64(v)) - 1)
		// 如果当前位置的数字大于0，我们就取反
		// 这里如果为负数，就说明我们这个数字出现过了
		if nums[pos] > 0 {
			nums[pos] = -nums[pos]
		}
	}
	// 下面我们只需要判断当前位置的数字是否为0就可以了
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
