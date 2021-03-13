// @Description 
// @Author 小游
// @Date 2021/03/13
package q34

// 解法1 使用双指针
func searchRange1(nums []int, target int) []int {
	if len(nums) ==0{
		return []int{-1,-1}
	}
	// 使用双指针
	low,high:=0,len(nums)-1
	// 当low大于high时我们就可以跳循环了
	for low < high {
		// 当high指针不为target时，我们的high指针--
		if nums[high] != target {
			high --
		}
		// 当low指针不为target时，我们的low指针++,这里我们需要同时判断
		if nums[low] != target {
			low ++
		}
		// 当两个指针都为target时我们就跳出循环
		if nums[high] == target && nums[low] == target {
			break
		}
	}
	// 只要有一个不是target那么就说明结果错误，我们就让两个指针同时为-1
	if nums[high] != target {
		high = -1
		low = -1
	}
	return []int{low,high}
}


func searchRange(nums []int, target int) []int {
	if len(nums) ==0{
		return []int{-1,-1}
	}
	low:=findLeft(nums,target)
	high:=findRight(nums,target)
	if low==len(nums) || nums[low]!= target {
		return []int{-1,-1}
	} else {
		return []int{low,high}
	}

}
// 找出第一次出现的位置
func findLeft(nums []int,target int) int {
	// 这里我们使用左闭右开的写法
	low,high,mid:=0,len(nums),0
	// 这里我们开始遍历
	for low < high {
		mid = (low+high) / 2
		// 关键部分在这里，为什么我们可以找到第一次出现的位置呢
		// 因为当我们的结果等于target时，我们就已经找到了
		// 这里我们让high等于mid，是因为我们的high就是正确答案，所以我们不能mid-1
		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
// 找出最后一次出现的位置
func findRight(nums []int,target int) int {
	low,high,mid:=0,len(nums),0
	// 这里我们开始遍历
	for low < high {
		mid = (low+high) / 2
		// 关键部分在这里，为什么我们可以找到第一次出现的位置呢
		// 因为当我们的结果大于target时，我们就相当于是找了了最后一次出现位置的后一位
		if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low-1
}