// @Description
// @Author 小游
// @Date 2021/03/25
package q504

import (
	"strconv"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	// 先判断是否为负数
	isNegative := num < 0
	// 如果是负数，直接取反
	if isNegative {
		num = -num
	}
	ans := ""
	// 下面这里就是计算7进制数了
	for num > 0 {
		a := num / 7
		b := num % 7
		ans = strconv.Itoa(b) + ans
		num = a
	}
	// 如果是负数那么就需要前面加一个负号
	if isNegative {
		ans = "-" + ans
	}
	return ans
}
