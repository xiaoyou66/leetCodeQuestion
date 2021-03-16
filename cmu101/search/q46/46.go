// @Description 
// @Author 小游
// @Date 2021/03/16
package q46

import "fmt"

func permute(nums []int) [][]int {
	ans:=make([][]int,0)
	backtracking(nums,0,ans)
	return ans
}

func backtracking(nums []int, level int, ans [][]int) {
	if level == len(nums) -1 {
		fmt.Println(nums)
		ans = append(ans, nums)
		return
	}
	for i := level; i < len(nums); i++ {
		nums[i],nums[level] = nums[level],nums[i]
		backtracking(nums,level+1,ans)
		nums[i],nums[level] = nums[level],nums[i]
	}
}