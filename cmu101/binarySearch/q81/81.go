// @Description 
// @Author 小游
// @Date 2021/03/13
package q81

func search(nums []int, target int) bool {
	for _,v:=range nums{
		if v == target {
			return true
		}
	}
	return false
}