// @Description
// @Author 小游
// @Date 2021/03/26
package q303

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	// 在我们开始阶段我分别统计前n位和并放入第n位数组上
	n := len(nums)
	sum := make([]int, n+1)
	sum[1] = nums[0]
	for i := 2; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	return NumArray{sum}
}

func (this *NumArray) SumRange(left int, right int) int {
	// 因为我们前面是已经计算好了的，所以我们这里直接两个和相减就是最终范围了
	return this.nums[right+1] - this.nums[left]
}
