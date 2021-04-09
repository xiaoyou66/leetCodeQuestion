// @Description
// @Author 小游
// @Date 2021/04/09
package q57

// 解法一 双指针
func twoSum1(nums []int, target int) []int {
	low, high := 0, len(nums)-1
	for low <= high {
		if nums[low]+nums[high] < target {
			low++
		} else if nums[low]+nums[high] > target {
			high--
		} else {
			return []int{nums[low], nums[high]}
		}
	}
	return []int{}
}

// 解法二 哈希表
func twoSum(nums []int, target int) []int {
	// 我们使用hash来暂存数据
	data := make(map[int]int)
	// 遍历数组求值
	for i := 0; i < len(nums); i++ {
		a := target - nums[i]
		// 判断在map中是否存在
		if _, ok := data[a]; ok {
			return []int{a, nums[i]}
		}
		// 把值存储进map中
		data[nums[i]] = i
	}
	return []int{}
}
