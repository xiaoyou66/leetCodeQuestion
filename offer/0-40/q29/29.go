// @Description
// @Author 小游
// @Date 2021/04/05
package q29

//func spiralOrder(matrix [][]int) []int {
//	if len(matrix) == 0 {
//		return []int{}
//	}
//	l,r,t,b,x:= 0,len(matrix[0])-1,0,len(matrix)-1,0
//	res:=make([]int,(r+1)*(b+1))
//	for  {
//		for i := l; i <= r; i++ { res[x] = matrix[t][i];x++ }
//		if t++; t > b { break }
//
//		for i := t; i <= b; i++ { res[x] = matrix[i][r];x++ }
//		if r--;l>r  { break }
//
//		for i := r; i >= l; i-- { res[x] = matrix[b][i];x++ }
//		if b--;t>b  { break }
//
//		for i := b; i >= t; i-- { res[x] = matrix[i][l];x++ }
//		if l++;l>r  { break }
//
//	}
//	return res
//}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	// 初始化上下左右
	t, b, l, r := 0, len(matrix)-1, 0, len(matrix[0])-1
	// 初始化结果数组
	res := make([]int, (r+1)*(b+1))
	x := 0
	// 开始进行遍历
	for {
		// 从左到右
		for i := l; i <= r; i++ {
			res[x] = matrix[t][i]
			x++
		}
		if t++; t > b {
			break
		}
		// 从上到下
		for i := t; i <= b; i++ {
			res[x] = matrix[i][r]
			x++
		}
		if r--; l > r {
			break
		}
		// 从右到左
		for i := r; i >= l; i-- {
			res[x] = matrix[b][i]
			x++
		}
		if b--; t > b {
			break
		}
		// 从下到上
		for i := b; i >= t; i-- {
			res[x] = matrix[i][l]
			x++
		}
		if l++; l > r {
			break
		}
	}
	return res
}
