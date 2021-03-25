// @Description
// @Author 小游
// @Date 2021/03/25
package q384

import (
	"math/rand"
)

type Solution struct {
	arr []int
}

func Constructor(nums []int) Solution {
	// 创建并返回solution结构体
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.arr
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.arr)
	res := make([]int, n)
	// 拷贝数据
	copy(res, this.arr)
	// 打乱数组
	for i := n - 1; i >= 0; i-- {
		// 返回返回[0, i]范围的整数
		j := rand.Intn(i + 1)
		// 交换元素的位置
		res[i], res[j] = res[j], res[i]
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
