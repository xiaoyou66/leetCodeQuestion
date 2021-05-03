// @Description
// @Author 小游
// @Date 2021/04/q01
package important

import (
	"fmt"
	"testing"
)

// 快排测试
func Test_quickSort(t *testing.T) {
	arr := []int{4, 8, 6, 45, 12, 787, 4, 5, 6, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 二分查找测试
func Test_binary(t *testing.T) {
	fmt.Println(binarySearch([]int{0, 5, 16, 20, 27, 30, 36, 44, 55, 60, 67, 71}, 55))
}

// 爬楼梯
func Test_comb(t *testing.T) {
	fmt.Println(climbStairs(4))
}

// 两数之和测试
func Test_twoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// 股票问题
func Test_profile(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

// 最大数组连续和
func Test_maxSub(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// 最长子串
func Test_longSub(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// 合并有序数组
func Test_merge(t *testing.T) {
	num1 := []int{1, 2, 3, 0, 0, 0}
	merge(num1, 3, []int{2, 5, 6}, 3)
	fmt.Println(num1)
}

// 全排列测试
func Test_permute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

// 全排列去重
func Test_recur(t *testing.T) {
	fmt.Println(permutation("aabc"))
}

// 三数之和测试
func Test_treeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
