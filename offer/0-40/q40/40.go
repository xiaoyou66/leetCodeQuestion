// @Description
// @Author 小游
// @Date 2021/04/07
package q40

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, arr[i])
	}
	return res
}
