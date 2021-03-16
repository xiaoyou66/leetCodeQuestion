// @Description 
// @Author 小游
// @Date 2021/03/16
package search

// 使用栈的写法
func maxAreaOfIsland1(grid [][]int) int {
	m:=len(grid)
	if m==0 {
		return 0
	}
	n:= len(grid[0])
	area, localArea :=0,0
	var x,y int
	// 遍历数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果当前节点为1，我们就可以进行遍历了
			if grid[i][j] ==1 {
				localArea = 1
				grid[i][j] = 0
				// 自己定义一个栈
				island:=Stack{}
				// 把当前值放入栈中
				island.Push([]int{i,j})
				// 遍历直到栈为空位置
				for !island.Empty() {
					// 获取当前栈顶的值
					r,c:=island.Pop()
					// 我们分别依次判断当前值的上下左右是否为空
					for k := 0; k < 4; k++ {
						// 这里我们使用了一个小技巧，每相邻两位即为上下左右四个方向之一
						x =r+ direction[k];y = c+direction[k+1]
						// 这里我们还需要判断一下x，y的范围是否在矩形内，以免越界
						if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
							// 这里我们就把grid置为0，然后把这个点放入栈中
							grid[x][y] = 0
							localArea++
							island.Push([]int{x,y})
						}
					}
				}
				if localArea > area {
					area= localArea
				}
			}
		}
	}
	return area
}
// 自己定义一个简单的栈
type Stack struct {
	i 	 int
	data [][]int
}
func (s *Stack) Push(k []int)  {
	s.data = append(s.data, k)
	s.i = len(s.data)-1
}
func (s *Stack) Pop() (x int,y int) {
	x = s.data[s.i][0]
	y = s.data[s.i][1]
	s.i--
	return
}
func (s *Stack) Empty() bool {
	return s.i < 0
}

var direction = []int{-1, 0, 1, 0, -1}
// 使用递归的写法
func maxAreaOfIsland(grid [][]int) int {
	// 当grid大小为0时，我们就退出循环
	if len(grid)==0 || len(grid[0])==0 {
		return 0
	}
	// 当前最大区域为0
	maxArea:=0
	// 我们开始遍历整个grid数组
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 当当前这个点为1的时候，我们就获取一下当前的区域信息
			if grid[i][j] == 1 {
				// 使用dfs来获取当前的地域信息
				area:=dfs(grid,i,j)
				// 更新最大的区域
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

// 使用一个辅函数，这个就是我们的关键部分了
func dfs(grid [][]int,r int,c int) int {
	// 因为是递归函数，所以我们需要设置一个递归的条件
	if grid[r][c] == 0 {
		return 0
	}
	// 如果r，c所在的值不为0，那么当前区域值就为1，同时把当前区域值置为0
	area:= 1
	var x,y int
	grid[r][c] = 0
	// 对当前位置进行遍历，判断上下左右四个方向
	for k := 0; k < 4; k++ {
		// 这里我们通过direction数组来实现获取当前位置的上下左右
		x=direction[k]+r
		y=direction[k+1]+c
		// 确保这个值在矩阵的范围内
		if x>=0 && y>=0 && x<len(grid) && y<len(grid[0]) {
			// 注意这我们不需要判断当前位置是否为0，因为dfs会自己计算，如果为0就会返回0
			area+=dfs(grid,x,y)
		}
	}
	return area
}