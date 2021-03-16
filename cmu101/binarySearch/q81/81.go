// @Description 
// @Author 小游
// @Date 2021/03/13
package q81

// 最笨的解法
func search1(nums []int, target int) bool {
	for _,v:=range nums{
		if v == target {
			return true
		}
	}
	return false
}

func search(nums []int, target int) bool {
	low,high:=0,len(nums)-1
	// 这里为什么要low等于high呢，因为我们当low等于high时我们还需要判断一下，确保没有漏掉值
	for low<=high {
		// 求中点
		mid:=(low+high)/2
		// 当我们找到时，直接返回
		if nums[mid] == target {
			return true
		} else if nums[low] == nums[mid] {
			// 当low和mid相等时，我们就尝试移动low，再判断一下
			low ++
		} else if nums[mid] <= nums[high]{
			// 当mid小于或等于high时，说明当前这个区间是递增的
			// 这里有两个细节，一个是我们计算时要确保target大于nums[mid]
			// 然后target小于等于nums[high](当等于时我们就必须移动low指针)
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			// 然后这里就是nums[mid] > nums[high]的情况了
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return false
}
