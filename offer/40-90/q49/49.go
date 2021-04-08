// @Description
// @Author 小游
// @Date 2021/04/08
package q49

func nthUglyNumber(n int) int {
	// 因为丑树是包含质因子2,3,5的数，所以我们就分别哟经abc来进行表示
	a := 0
	b := 0
	c := 0
	// 初始化dp数组
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		// 我们分别求一下每个质因子和较小的数相乘得到的结果
		an := dp[a] * 2
		bn := dp[b] * 3
		cn := dp[c] * 5
		// 当前的丑数值就应该上面三个的最小值
		dp[i] = min(an, min(bn, cn))
		// 判断一下是哪个最小，然后我们更新一下下标
		// 因为我们要确保丑数要每个质因子都要乘一下，所以我们采用3个下标的方式
		if dp[i] == an {
			a++
		}
		if dp[i] == bn {
			b++
		}
		if dp[i] == cn {
			c++
		}
	}
	return dp[n-1]
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
