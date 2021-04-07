// @Description
// @Author 小游
// @Date 2021/04/07
package q44

import "strconv"

func findNthDigit(n int) int {
	// digit表示数位
	// start表示起始数字
	// count表示数位的数量
	digit := 1
	start := 1
	count := 9
	// 当n小于数位数量的时候，我们就退出循环
	for n > count {
		// 这里是根据我们的数学规律进行计算得到的
		n -= count
		digit += 1
		start *= 10
		count = digit * start * 9
	}
	// num表示所求数位在第几个数字中
	num := start + (n-1)/digit
	// 然后这里确定我们的数位
	return int(strconv.Itoa(num)[(n-1)%digit] - '0')
}
