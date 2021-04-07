// @Description
// @Author 小游
// @Date 2021/03/31
package q15

func hammingWeight(num uint32) int {
	res := 0
	// 这题其实
	for num > 0 {
		// 使用位运算，如果最低位为1那么和1相与后就为1，否则就为0
		res += int(num & 1)
		// 把数字往左移
		num >>= 1
	}
	return res
}
