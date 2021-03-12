// @Description 
// @Author 小游
// @Date 2021/03/12
package q167

// 双指针算法
func twoSum(numbers []int, target int) []int {
	// 这里我们使用两个指针，一个执行0，一个指向数组的最高端
	low,high:=0,len(numbers)-1
	// 当low指针大于high指针的时候我们就可以退出循环了
	// （我们没有必要两边都走一遍，因为当low等于high的时候，我们就已经把整个数组遍历完了）
	for low < high {
		// 当两个指针的值大于我们的目标时，说明值大了，我们就需要把high指针减小
		sum := numbers[low] + numbers[high]
		if sum == target{
			break
		} else if sum  > target {
			high --
		} else if sum < target { // 当两个指针的值小于我们的目标时，说明值小了，我们需要加大low指针
			low ++
		}
	}
	// 上面我们已经计算完了，直接返回数据
	return []int{low+1,high+1}
}