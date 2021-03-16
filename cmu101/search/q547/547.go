// @Description 
// @Author 小游
// @Date 2021/03/16
package q547

func findCircleNum(isConnected [][]int) int {
	n:=len(isConnected)
	count:=0
	visited:=make([]bool,n)
	for i := 0; i < n; i++ {
		if !visited[i]  {
			dfs(isConnected,i,visited)
			count++
		}
	}
	return count
}

func dfs(area [][]int,i int,visited []bool) {
	visited[i] = true
	for k:=0;k<len(area);k++{
		if area[i][k] == 1 && !visited[k] {
			dfs(area,k,visited)
		}
	}
}
