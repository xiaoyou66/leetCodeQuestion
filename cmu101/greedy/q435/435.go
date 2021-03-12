// @Description https://leetcode-cn.com/problems/non-overlapping-intervals/
// @Author 小游
// @Date 2021/03/12
package q435

import "sort"

// 这里我们为了进行简单排序，需要重写sort方法
type Intervals [][]int
// 获取长度
func (a Intervals) Len() int {
	return len(a)
}
// 交换数据
func (a Intervals) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
// 重写less方法，这里是比较两个下标
func (a Intervals) Less(i, j int) bool {
	// 我们这里是在进行区间比较，只要发现开头或结尾i比j小，那么就是小的，否则就是大的
	for k := 0; k < len(a[i]); k++ {
		if a[i][k] < a[j][k] {
			return true
		} else if a[i][k] == a[j][k] {
			continue
		} else {
			return false
		}
	}
	return true
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 这里是对我们传入的参数进行区间排序
	// 这里排序完后，那些开头比较小的会排在前面
	// 然后如果开头相等，但是结尾较小的也会在前面
	sort.Sort(Intervals(intervals))
	// 这里因为我们要让整个区间连续不能重叠
	// 所以我们使用pre来表示当前已经连续了的区间最小位置
	pre, res := 0, 1
	// 然后我们就开始进行比较
	for i := 1; i < len(intervals); i++ {
		// 如果当前区间的开头比 pre区间的结尾大或者相等
		if intervals[i][0] >= intervals[pre][1] {
			// 这里表示区间不重叠，此时我们就可以让pre等于当前的区间，res表示连续的区间数
			res++
			pre = i
		// 如果当前区间的结尾比pre更小，那么我们可以把pre换成i
		// 这里说一下这一步的原因，因为我们要确保每次都选结尾最早的
		// 此时我们的数组是排好序了的，pre的开头一定会小于或者等于i的开头
		// 这个时候我们直接换成i可以减小结尾的长度
		} else if intervals[i][1] < intervals[pre][1] {
			pre = i
		}
	}
	return len(intervals) - res
}