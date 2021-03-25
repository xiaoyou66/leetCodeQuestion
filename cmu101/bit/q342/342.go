// @Description
// @Author 小游
// @Date 2021/03/25
package q342

func isPowerOfFour(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && (n&1431655765) > 0
}
