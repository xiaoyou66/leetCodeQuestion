// @Description https://leetcode-cn.com/problems/two-sum/submissions/
// @Author 小游
// @Date 2021/03/06
package q001

func twoSum(nums []int, target int) []int {
	// 创建一个map
	m := make(map[int]int)
	// 遍历整个nums
	for i := 0; i < len(nums); i++ {
		// 关键部分，我没有想到其实我们知道了结果，
		// 知道了其中一个数，其实就知道了另一个数
		another := target - nums[i]
		// 我们的map key存储的是值，value存储的是位置
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		// 如果在map中不存在的话，那我们就把这个数字存入map中
		m[nums[i]] = i
	}
	return nil
}

//func twoSum(nums []int,target int) []int {
//	// 初始化一个map
//	m:=new(map[int]int)
//	// 遍历nums数组
//
//}
