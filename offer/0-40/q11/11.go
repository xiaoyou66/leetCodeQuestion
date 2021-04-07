// @Description
// @Author 小游
// @Date 2021/03/30
package q11

import (
	"math"
)

// 解法一 暴力计算
func minArray1(numbers []int) int {
	// 最简单的方法，遍历数组
	min := math.MaxInt32
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return min
}

// 解法2 二分查找
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		mid := low + (high-low)/2
		// 开始判断比较
		if numbers[mid] > numbers[high] {
			// 关键点一 low等于mid+1，自己画图证明
			low = mid + 1
		} else if numbers[mid] < numbers[high] {
			// 关键点二 high=mid
			high = mid
		} else {
			high--
		}
	}
	return numbers[low]
}
