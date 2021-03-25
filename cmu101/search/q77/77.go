// @Description
// @Author 小游
// @Date 2021/03/17
package q77

func combine(n int, k int) [][]int {
	// ans表示结果，comb表示当前的组合，count表示组合的位数
	ans := make([][]int, 0)
	comb := make([]int, k)
	count := 0
	// 进行回溯获取结果，题目是从1到n所以我们的pos一开始为1
	getOrder(&ans, comb, count, 1, n, k)
	return ans
}
func getOrder(ans *[][]int, comb []int, count int, pos int, n int, k int) {
	// 当count等于k的时候，我们就可以把答案加到结果里去了
	if count == k {
		tmp := make([]int, k)
		copy(tmp, comb)
		*ans = append(*ans, tmp)
		return
	}
	// 因为题目是1到n，所以可以为n，然后我们从pos开始进行计算
	for i := pos; i <= n; i++ {
		// count表示当前组合的数量,我们把当前的位置放入数组中
		comb[count] = i
		count++
		// 这里我们进行计算获取ans结果
		// 这个地方可能难以理解，为什么我们这里可以实现所有组合呢
		// 就拿最简单的两个来说，第一个位置我们可以取 pos -> n
		// 第二个位置同样会执行这个递归 这里会取 pos+1 -> n
		getOrder(ans, comb, count, i+1, n, k)
		count-- // 回溯，返回上一次状态
	}
}
