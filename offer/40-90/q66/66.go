// @Description
// @Author 小游
// @Date 2021/04/10
package q66

func constructArr(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	// 初始化数组
	b := make([]int, len(a))
	// 先把b[0]置为1
	b[0] = 1

	tmp := 1
	// 这里是计算下三角，
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	// 这里我们计算上三角，然后计算得到的tmp去乘
	for i := len(a) - 2; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}
	return b
}
