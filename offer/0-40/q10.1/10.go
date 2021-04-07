// @Description
// @Author 小游
// @Date 2021/03/30
package q10_1

// 使用动态规划进行求解
func fib(n int) int {
	// 这题使用动态规划
	// 我们让a为n-2，然后b表示n-1
	//(其实一开始a表示0 b表示1，所以我们最后计算的时候直接输出a就可以了)
	a := 0
	b := 1
	sum := 0
	for i := 0; i < n; i++ {
		// 因为题目规定要取模计算，所以我们进行取模
		sum = (a + b) % (1e9 + 7)
		// 然后更新一下(n-2 和 n-1)
		a = b
		b = sum
	}
	// 因为我们要计算第n项
	return a
}
