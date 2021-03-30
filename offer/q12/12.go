// @Description
// @Author 小游
// @Date 2021/03/30
package q12

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	// 判断大小
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i int, j int, k int, word string) bool {
	// 判断i，j是否合法
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	// 判断当前位数是否等于字符串的位数
	if k == len(word)-1 {
		return true
	}
	// 我们把当前这位置为0表示我们访问过了
	board[i][j] = 0
	// 然后我们进行上下左右遍历
	res := dfs(board, i-1, j, k+1, word) || dfs(board, i, j-1, k+1, word) || dfs(board, i+1, j, k+1, word) || dfs(board, i, j+1, k+1, word)
	// 没有的话，我们在进行回溯
	board[i][j] = word[k]
	return res
}
