// @Description
// @Author 小游
// @Date 2021/03/17
package q70

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 1, 1, 1
	for i := 2; i <= n; i++ {
		a = b + c
		c = b
		b = a
	}
	return a
}
