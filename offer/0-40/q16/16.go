// @Description
// @Author 小游
// @Date 2021/03/31
package q16

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	var res = float64(1)
	// 当n为负数时,我们就翻转一下
	if n < 0 {
		x = 1 / x
		n = -n
	}
	// 进行位运算操作
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}
