// @Description
// @Author 小游
// @Date 2021/04/01
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

// 两数之和测试
func Test_twoSum(t *testing.T) {
	fmt.Println(binarySearch([]int{2, 7, 11, 15}, 9))
}

// 全排列测试
func Test_permute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

// 三数之和测试
func Test_treeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
