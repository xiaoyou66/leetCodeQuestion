// @Description https://leetcode-cn.com/problems/sqrtx/
// @Author 小游
// @Date 2021/03/13
package q69


func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	// 定义一个left和right两个指针，以及最后的答案
	left, right, res := 1, x, 0
	for left <= right {
		// 右移一位相当于除2操作，右移n位相当于除以2的n次方
		// 我们这里修改一下right
		// 我们这里相当于求解 f (x) = x^2 − a = 0
		// 这里这样做可以让我们的值更加平均
		mid := left + ((right - left) >> 1)
		// 小于时我们移动left指针，同时记录当前结果
		// 我们的结果一般都是使用左指针来
		if mid < x/mid {
			left = mid + 1
			res = mid
		} else if mid == x/mid {
			return mid
		} else {
			right = mid - 1
		}
	}
	return res
}
