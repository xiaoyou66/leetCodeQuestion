// @Description
// @Author 小游
// @Date 2021/04/10
package q64

var res = 0
var x bool

// 计算
func sumNums(n int) int {
	// 计算时我们需要把res置为0，避免
	res = 0
	return count(n)
}

// 开始计算
func count(n int) int {
	x = n > 1 && sumNums(n-1) > 0
	res += n
	return res
}
