// @Description
// @Author 小游
// @Date 2021/03/26
package q739

type Stack struct {
	data []int
}

func (s *Stack) Push(k int) {
	s.data = append(s.data, k)
}
func (s *Stack) Pop() (x int) {
	x = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return
}
func (s *Stack) Empty() bool {
	return len(s.data) == 0
}
func (s *Stack) Top() int {
	return s.data[len(s.data)-1]
}
func dailyTemperatures(T []int) []int {
	n := len(T)
	ans := make([]int, n)
	// 维护一个单调栈来表示每天的温度
	indices := Stack{}
	// 遍历
	for i := 0; i < n; i++ {
		// 判断indices是否为空
		for !indices.Empty() {
			// 获取栈顶为第几天
			pre := indices.Top()
			// 判断一下当天的温度是否比栈顶的高
			if T[i] <= T[pre] {
				break
			}
			// 如果当天气温比栈顶的高，那么我们就可以知道要等多少天数可以观察到更高温的天气了
			indices.Pop()
			// 这里我们就知道结果了
			ans[pre] = i - pre
		}
		// 把当天放入栈中
		indices.Push(i)
	}
	return ans
}
