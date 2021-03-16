// @Description 
// @Author 小游
// @Date 2021/03/16
package q417


var direction = []int{-1, 0, 1, 0, -1}

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	m,n:=len(matrix),len(matrix[0])
	var ans [][]int
	reachP:=make([][]bool,m)
	reachA:=make([][]bool,m)
	for i := 0; i < m; i++ {
		reachP[i] = make([]bool,n)
		reachA[i] = make([]bool,n)
	}
	for i := 0; i < m; i++ {
		dfs(matrix,reachP,i,0)
		dfs(matrix,reachA,i,n-1)
	}
	for i := 0; i < n; i++ {
		dfs(matrix,reachP,0,i)
		dfs(matrix,reachA,m-1,i)
	}

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if reachA[i][j] && reachP[i][j] {
				ans = append(ans, []int{i,j})
			}
		}
	}
	return ans
}

func dfs(matrix [][]int,reach [][]bool,r int,c int) {
	if reach[r][c]{
		return
	}
	reach[r][c] = true
	var x,y int
	for i := 0; i < 4; i++ {
		x = direction[i] + r
		y = direction[i+1] + c
		if x>=0 && y>=0 && x < len(matrix) && y<len(matrix[0]) && matrix[r][c] <= matrix[x][y] {
			dfs(matrix,reach,x,y)
		}
	}
}