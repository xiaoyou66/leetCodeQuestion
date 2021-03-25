// @Description
// @Author 小游
// @Date 2021/03/25
package q72

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// 创建dp数组一维表示第一个字符串，二维表示第二个字符串
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 准备计算
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			// 当i等于0时，表示第一个字符串没有内容
			// 所以我们如果想转换为第二个字符串时，需要操作j次（此时第二个字符串长度为j）
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				// 当j等于0时，表示第二个字符串为空
				// 如果我们需要转为第二个字符串，需要操作i次（就是删除i次，第一个字符串长度为i）
				dp[i][j] = i
			} else {
				var eq int
				// 判断当前位置的字符串是否相同(注意，数组的长度从0开始，所以我们要-1)
				if word1[i-1] == word2[j-1] {
					eq = 0
				} else {
					eq = 1
				}
				// 这里我们需要考虑三种情况
				// 1.当我们选择修改时，需要修改的次数是dp[i-1][j-1]（这个自己理解吧）
				// 2.当我们选择插入 j 位置/删除 i 位置时，需要修改的次数为 dp[i-1][j] + 1
				// 3.当我们选择插入 i 位置/删除 j 位置时，需要修改的次数为 dp[i][j-1] + 1
				dp[i][j] = min(dp[i-1][j-1]+eq, min(dp[i-1][j]+1, dp[i][j-1]+1))
			}
		}
	}
	return dp[m][n]
}
func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
