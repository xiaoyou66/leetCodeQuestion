// @Description
// @Author 小游
// @Date 2021/04/09
package q56

func singleNumbers1(nums []int) []int {
	// 最终结果
	var res []int
	data := make(map[int]int)
	// 第一遍，统计出现的次数
	for _, v := range nums {
		data[v]++
	}
	// 第二遍找出只出现一次的数
	for k, v := range data {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

// 解法二 分组异或
func singleNumbers(nums []int) []int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	div := 1
	for div&ret == 0 {
		div <<= 1
	}
	a, b := 0, 0
	for _, v := range nums {
		if (div & v) != 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}
