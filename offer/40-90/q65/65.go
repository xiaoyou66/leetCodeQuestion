// @Description
// @Author 小游
// @Date 2021/04/10
package q65

func add(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1
		a ^= b
		b = c
	}
	return a
}
