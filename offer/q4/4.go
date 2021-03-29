// @Description
// @Author 小游
// @Date 2021/03/29
package q4

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 简单过滤特殊情况
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n := len(matrix[0])
	// i和j分别表示两个下标
	i := len(matrix) - 1
	j := 0
	// for循环遍历
	for i >= 0 && j < n {
		// 判断当前这个点是否大于目标值，如果大于我们就需要缩小i
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] == target {
			return true
		} else {
			// 否则我们就需要增加j的值
			j++
		}
	}
	return false
}
