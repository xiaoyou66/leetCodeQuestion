// @Description
// @Author 小游
// @Date 2021/04/07
package q46

import "strconv"

func translateNum(num int) int {
	// 为了方便计算我们先把数字转换为字符串
	s := strconv.Itoa(num)
	// 然后使用a，b来暂存,a表示i-2 b表示i-1
	a := 1
	b := 1
	// 然后我们开始计算
	for i := 2; i <= len(s); i++ {
		// 每次截取两个字符
		tmp := s[i-2 : i]
		var c int
		// 然后比较一下大小
		if tmp >= "10" && tmp <= "25" {
			c = a + b
		} else {
			c = a
		}
		b = a
		a = c
	}
	return a
}
