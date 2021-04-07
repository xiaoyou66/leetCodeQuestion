// @Description
// @Author 小游
// @Date 2021/04/07
package q39

import "sort"

// 解法一 使用map统计
func majorityElement1(nums []int) int {
	// 统计每位出现的次数
	count := make(map[int]int)
	max := 0
	// 遍历求值
	for _, v := range nums {
		count[v]++
		if count[v] > count[max] {
			max = v
		}
	}
	return max
}

// 解法二 排序
func majorityElement2(nums []int) int {
	// 先排序
	sort.Ints(nums)
	return nums[(0+len(nums))/2]
}

// 解法三 位运算
func majorityElement3(nums []int) int {
	major := 0
	n := len(nums)
	mask := 1
	for i := 0; i < 32; i++ {
		bitCounts := 0
		for j := 0; j < n; j++ {
			if nums[j]&mask == 1 {
				bitCounts++
			}
			if bitCounts > n/2 {
				major |= mask
				break
			}
		}
		mask <<= 1
	}
	return major
}

// 解法四 摩尔投票法
func majorityElement(nums []int) int {
	// 一开始我们让这个人表示第一个国家
	major := nums[0]
	// 国家人数设置为1
	count := 1
	// 下面我们开始相互对抗
	for i := 1; i < len(nums); i++ {
		// 如果国家人数为0，那么我们就选一个新的国家
		if count == 0 {
			major = nums[i]
			count++
		} else if nums[i] == major {
			// 如果数字相同，那么国家人数+1
			count++
		} else {
			// 否则的话，我们国家人数就要派一个人同归于尽，人数少一
			count--
		}
	}
	// 最后还剩下的国家就是大于一半的情况了
	return major
}
