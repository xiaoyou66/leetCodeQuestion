// @Description
// @Author 小游
// @Date 2021/04/09
package q53

func search(nums []int, target int) int {
	// 我们使用二分查找查找数据
	low := 0
	high := len(nums) - 1
	count := 0
	// 当low大于high时退出循环
	for low <= high {
		// 计算中点
		mid := low + (high-low)/2
		// 下面这个是二分查找的步骤，我这里就不多说了
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			// 当我们找到是，因为我们要统计出现的次数
			t := mid
			count = 1
			// 所以我们需要向前和向后遍历
			t++
			mid--
			for t < len(nums) && nums[t] == target {
				count++
				t++
			}
			for mid >= 0 && nums[mid] == target {
				count++
				mid--
			}
			return count
		}
	}
	return count
}
