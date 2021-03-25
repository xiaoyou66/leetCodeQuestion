// @Description 公倍数与公因数
// @Author 小游
// @Date 2021/03/25
package gcd

// 求解最大公因数
func gcd(a int, b int) int {
	// 这里我们使用辗转相除法来求解
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// 求解最小公倍数(只需要把两个数相乘然后再除以最大公因数就可以得到最小公倍数了)
func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}
