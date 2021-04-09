// @Description
// @Author 小游
// @Date 2021/04/09
package q57_2

// 解法1 暴力
func findContinuousSequence1(target int) [][]int {
	// res表示最最终结果
	var res [][]int
	// sum表示总数，然后limit表示我们从哪里结束（因为答案是至少两个连续序列，所以我们这里取一半就可以了）
	sum, limit := 0, (target)/2
	// 遍历
	for i := 1; i <= limit; i++ {
		for j := i; ; j++ {
			sum += j
			// 当总和大于我们的目标值时，退出循环
			if sum > target {
				sum = 0
				break
			} else if sum == target {
				// 当sum等于我们的目标时，我们把结果加入数组中
				tmp := make([]int, j-i+1)
				for k := i; k <= j; k++ {
					tmp[k-i] = k
				}
				// 我们添加当前结果
				res = append(res, tmp)
				sum = 0
				break
			}
		}
	}
	return res
}

// 解法二 双指针
func findContinuousSequence2(target int) [][]int {
	var res [][]int
	r := 2
	for l := 1; l < r; {
		// 区间求和公式
		sum := (l + r) * (r - l + 1) / 2
		// 当sum等于target时，此时返回结果
		if sum == target {
			tmp := make([]int, r-l+1)
			for i := l; i <= r; i++ {
				tmp[i-l] = i
			}
			res = append(res, tmp)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return res
}

// 解法三 滑动窗口
func findContinuousSequence(target int) [][]int {
	low := 1
	high := 2
	s := 3
	var res [][]int
	for low < high {
		if s == target {
			tmp := make([]int, high-low+1)
			for k := low; k <= high; k++ {
				tmp[k-low] = k
			}
			res = append(res, tmp)
		}
		if s >= target {
			s -= low
			low++
		} else {
			high++
			s += high
		}
	}
	return res
}
