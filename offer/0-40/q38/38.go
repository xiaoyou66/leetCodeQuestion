// @Description
// @Author 小游
// @Date 2021/03/31
package q38

func permutation(s string) []string {
	var data []string
	//visited := make([]bool, len(s))
	dfs(&data, 0, []byte(s))
	return data
}

func dfs(data *[]string, k int, s []byte) {
	if k == len(s)-1 {
		tmp := make([]byte, len(s))
		copy(tmp, s)
		*data = append(*data, string(tmp))
		return
	}
	set := map[byte]bool{}
	// 开始遍历
	for i := k; i < len(s); i++ {
		// 字符串不存在我们才进行交换
		if _, ok := set[s[i]]; !ok {
			set[s[i]] = true
			s[k], s[i] = s[i], s[k]
			dfs(data, k+1, s)
			s[k], s[i] = s[i], s[k]
		}
	}
}
