// @Description
// @Author 小游
// @Date 2021/03/25
package q461

func hammingDistance(x int, y int) int {
	// 进行异或运算，相同为0，不同为1
	diff := x ^ y
	ans := 0
	// 当我们的diff为0时我们就计算完了
	for diff > 0 {
		// 当我们这位不相同时，我们异或运算的结果就为1
		ans += diff & 1
		// 然后我们往右移动一位
		diff >>= 1
	}
	return ans
}
