// @Description
// @Author 小游
// @Date 2021/03/19
package q38

func permutation(s string) []string {
	var ans []string
	bfs(s, 0, &ans)
	return ans
}

func bfs(s string, level int, ans *[]string) {
	if level == len(s)-1 {
		*ans = append(*ans, s)
		return
	}
}
