// @Description
// @Author 小游
// @Date 2021/04/01
package important

import (
	"math"
	"sort"
)

// 二分查找
func binarySearch(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			return true
		}
	}
	return false
}

// 爬楼梯问题
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	// a表示第一阶，b表示第二阶，后面a就表示n-2 b表示n-1
	a := 1
	b := 2
	// sum就表示总数
	sum := 0
	// 然后我们就可以计算每一阶的了
	for i := 3; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return sum
}

// 两数之和
func twoSum(nums []int, target int) []int {
	// 我们使用hash来存储数据
	hash := make(map[int]int, len(nums))
	// 遍历整个数组
	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		// 判断这个值是否在map中
		if _, ok := hash[tmp]; ok {
			return []int{tmp, i}
		} else {
			hash[i] = 1
		}
	}
	return []int{}
}

// 股票问题
func maxProfit(prices []int) int {
	// 如果长度只是1，那么我们就返回0
	if len(prices) <= 1 {
		return 0
	}
	// 记录一个最小值和最大利润
	minPrice := math.MaxInt32
	maxProfile := 0
	// 开始遍历数组进行计算
	for i := 0; i < len(prices); i++ {
		// 先找到一个最小价格买入
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfile {
			// 然后我们以最大价格卖出
			maxProfile = prices[i] - minPrice
		}
	}
	return maxProfile
}

// 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 使用i来表示当前数组的结尾
	i := n + m - 1
	// 我们从后往前进行遍历
	for m >= 0 && n >= 0 {
		if nums1[m] > nums1[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}
	// 判断n是否排序完毕
	for n >= 0 {
		nums1[i] = nums2[n]
		i--
		n--
	}
}

// 连续子数组的最大和
func maxSubArray(nums []int) int {
	// 默认情况下，我们的最大值就是第0号
	res := nums[0]
	// 下面我们开始顺序遍历
	for i := 0; i < len(nums); i++ {
		// 这个地方是关键，看自己理解了
		nums[i] += max(nums[i-1], 0)
		// 最后我们求一个最大值
		res = max(res, nums[i])
	}
	return res
}

// 求最大值
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 解法一，滑动窗口
func lengthOfLongestSubstring(s string) int {
	// 使用一个map来存储位置信息
	data := make(map[byte]int)
	// n表示s的长度
	n := len(s)
	// 声明一个右指针表示起始第一个，然后ans表示最大的长度
	rk, ans := 0, 0
	for i := 0; i < n; i++ {
		if n != 0 {
			// 我们尝试删除前一个元素，然后再进行遍历判断
			delete(data, s[i-1])
		}
		// 如果右指针小于s的长度，并且当前位置也没重复的话,我们就移动一下右指针
		for rk < n && data[s[rk]] == 0 {
			data[s[rk]]++
			rk++
		}
		// 最后我们就可以统计一下长度了
		ans = max(ans, rk-i)
	}
	return ans
}

// 全排列问题
func permute(nums []int) [][]int {
	var data [][]int
	getOrder(nums, 0, &data)
	return data
}

// 我们需要使用一个辅助数组来存储数据
func getOrder(nums []int, index int, ans *[][]int) {
	// 当我们的index等于nums时，结束循环
	if index == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*ans = append(*ans, tmp)
		return
	}
	// 然后我们开始进行排序
	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		getOrder(nums, index+1, ans)
		nums[index], nums[i] = nums[i], nums[index]
	}
}

// 三数之和
func threeSum(nums []int) [][]int {
	// 数组的长度
	n := len(nums)
	// 给数组进行排序
	sort.Ints(nums)
	// 最后的排序结果
	ans := make([][]int, 0)
	// 开始进行遍历,这个是第一个指针
	for k := 0; k < n-2; k++ {
		// 首先判断k是否小于或者等于0，如果不是我们就跳出循环
		// 因为我们的数组实际上是已经排好序了的，所以我们k一定是小于或者等于0的
		if nums[k] > 0 {
			break
		}
		// 如果遇到相等的，我们可以先跳过，因为相等的我们已经计算过了，不需要重复计算
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		// i表示第二个指针，j表示第三个指针
		i := k + 1
		j := n - 1
		for i < j {
			// 先计算一下和
			sum := nums[k] + nums[i] + nums[j]
			// 如果当前小于0，我们可以移动一下i指针，当然我们还需要考虑指针相等的情况
			if sum < 0 {
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				// 我们无论如何都要+1操作，因为上面那个for循环可能不执行
				i++
			} else if sum > 0 {
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				j--
			} else {
				ans = append(ans, []int{nums[k], nums[i], nums[j]})
				// 我们继续寻找下一个
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			}
		}
	}
	return ans
}
