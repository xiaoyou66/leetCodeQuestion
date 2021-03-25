// @Description
// @Author 小游
// @Date 2021/03/16
package q46

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	// 直接调用辅助函数来获取值即可
	backtracking(nums, 0, &ans)
	return ans
}

// 注意这里为了能传递值，我们必须使用指针
func backtracking(nums []int, level int, ans *[][]int) {
	// 当我们的level为nums时，就说明已经到最后一位了
	if level == len(nums)-1 {
		// 因为nums是个指针，我们必须使用临时变量拷贝值，要不然会报错
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		// 我们直接把nums添加到结果里就行了
		*ans = append(*ans, tmp)
		return
	}
	// 这里就是关键部代码了
	for i := level; i < len(nums); i++ {
		// 这里其实就是i和当前位置换一下值（这个i是不断递增的）
		nums[i], nums[level] = nums[level], nums[i]
		// 下面这里就是进行替换和判断
		backtracking(nums, level+1, ans)
		// 替换回去，进行下一轮交换
		nums[i], nums[level] = nums[level], nums[i]
	}
}
