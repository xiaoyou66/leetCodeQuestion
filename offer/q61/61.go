// @Description
// @Author 小游
// @Date 2021/04/09
package q61

import "sort"

func isStraight1(nums []int) bool {
	repeat := make(map[int]bool)
	max := 0
	min := 14
	for _, v := range nums {
		// 跳过大小王
		if v == 0 {
			continue
		}
		// 更新一下最大牌和最小牌
		max = Max(max, v)
		min = Min(min, v)
		// 判断是否有重复元素
		if _, ok := repeat[v]; ok {
			return false
		}
		// 把当前排加入set中
		repeat[v] = true
	}
	// 判断是否可以构成顺子
	return max-min < 5
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 解法二 排序+遍历
func isStraight(nums []int) bool {
	joker := 0
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}
