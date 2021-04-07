// @Description
// @Author 小游
// @Date 2021/04/07
package q43

func countDigitOne(n int) int {
	digit := 1
	res := 0
	high := n / 10
	cur := n % 10
	low := 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res += high * digit
		} else if cur == 1 {
			res += high*digit + low + 1
		} else {
			res += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return res
}
