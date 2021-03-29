// @Description https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// @Author 小游
// @Date 2021/03/29
package q3

import (
	"sort"
)

// 解法一，最简单的使用map来进行判断
func findRepeatNumber1(nums []int) int {
	// 直接使用map
	count := make(map[int]int)
	for _, v := range nums {
		if _, ok := count[v]; ok {
			return v
		} else {
			count[v] = 1
		}
	}
	return -1
}

// 解法二，使用排序
func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)
	for k := range nums {
		if nums[k] == nums[k+1] {
			return nums[k]
		}
	}
	return -1
}

// 解法三，原地置换
func findRepeatNumber(nums []int) int {
	// 特殊判断,这里我们要保证满足条件
	if nums == nil || len(nums) == 0 {
		return -1
	}
	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return -1
		}
	}
	// 遍历数组
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			// 交换这两个元素的位置，这个地方是关键，我们可以这样理解，就是如果数组里面有重复的元素
			// 那么nums[i]就必然会有多个相同的值，这里我们交换是为了判断
			// 比如我们这里nums[i] = 1，i=2，此时我们的nums[1]=2
			// 下面我们如果遇到一个 j=3 如果nums[j]=1 时，此时
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
