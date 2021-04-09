// @Description
// @Author 小游
// @Date 2021/04/09
package q53_2

func missingNumber1(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == mid {
			low = mid + 1
		} else if nums[mid] > mid {
			// 如果前一个是相等的，那么缺失的就是当前这个啦
			if mid != 0 && nums[mid-1] == mid-1 {
				return mid
			}
			// 否则我们再缩小一下范围
			high = mid - 1
		}
	}
	// 如果没找到，那么要么是缺失首部，要么是缺失尾部
	if nums[0] != 0 {
		return 0
	} else {
		return len(nums)
	}
}

func missingNumber(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
