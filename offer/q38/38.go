// @Description
// @Author 小游
// @Date 2021/03/31
package q38

func permutation(s string) []string {
	var data []string
	visited := make([]bool, len(s))
	dfs(&data, []byte{}, visited, 0, s)
	return data
}

func dfs(data *[]string, current []byte, visited []bool, k int, s string) {
	if k == len(s) {
		//fmt.Println(string(current))
		*data = append(*data, string(current))
		return
	}
	for i := 0; i < len(s); i++ {
		// 判断是否访问过了
		if !visited[i] {
			visited[i] = true
			dfs(data, append(current, s[i]), visited, k+1, s)
			visited[i] = false
		}
	}
}
