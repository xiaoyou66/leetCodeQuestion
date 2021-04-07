// @Description
// @Author 小游
// @Date 2021/04/05
package q21

// 解法一 双指针
func exchange1(nums []int) []int {
	low, high := 0, len(nums)-1
	for low < high {
		// 首先我们找到一个偶数
		for high > low && nums[high]%2 == 0 {
			high--
		}
		// 然后找到一个奇数
		for high > low && nums[low]%2 != 0 {
			low++
		}
		// 然后我们进行替换
		nums[low], nums[high] = nums[high], nums[low]
		// 指针移动
		low++
		high--
	}
	return nums
}

// 解法二 快慢指针
func exchange(nums []int) []int {
	// 定义快慢两个指针
	slow, fast := 0, 0
	// 遍历直到fast遍历到数组尾巴为止
	for fast < len(nums) {
		// 当我们发送fast为奇数时，我们交换一下内容
		if nums[fast]&1 > 0 {
			// 因为slow是偶数（fast只会交换奇数，交换后的就是偶数了）
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		// 移动快指针
		fast++
	}
	return nums
}
