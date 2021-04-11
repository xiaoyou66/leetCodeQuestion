// @Description
// @Author 小游
// @Date 2021/04/10
package q62

func lastRemaining(n int, m int) int {
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}
