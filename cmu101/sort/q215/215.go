// @Description 
// @Author 小游
// @Date 2021/03/15
package q215

func findKthLargest(nums []int, k int) int {
	// 先对数组进行排序
	quickSort(nums,0,len(nums)-1)
	return nums[len(nums)-k]
}

// 快速排序找到中点
func finMid(arr []int,low int,high int) int  {
	// 我们就使用第一个作为临时点
	tmp:=arr[low]
	for low < high {
		// 错误一：这里必须要low<high
		for low < high && arr[high] >= tmp {
			high --
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= tmp {
			low ++
		}
		arr[high] = arr[low]
	}
	arr[low] = tmp
	return low
}

// 快速排序
func quickSort(arr []int,low int,high int)  {
	//错误2 这里是要确保low<high
	if low < high {
		mid:= finMid(arr,low,high)
		quickSort(arr,low,mid-1)
		quickSort(arr,mid+1,high)
	}
}