// @Description
// @Author 小游
// @Date 2021/03/26
package q1

func twoSum(nums []int, target int) []int {
	// 创建map并遍历数组
	arr := map[int]int{}
	for k, v := range nums {
		arr[v] = k
	}
	// 求解
	for k, v := range nums {
		a := target - v
		if index, ok := arr[a]; ok && index != k {
			return []int{k, index}
		}
	}
	return []int{}
}
