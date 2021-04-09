// @Description
// @Author 小游
// @Date 2021/04/09
package q51

func reversePairs(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	cop := make([]int, n)
	copy(cop, nums)
	tmp := make([]int, n)
	return reversePairs1(cop, 0, n-1, tmp)
}

func reversePairs1(nums []int, left int, right int, tmp []int) int {
	if left == right {
		return 0
	}
	mid := left + (right-left)/2
	leftPairs := reversePairs1(nums, left, mid, tmp)
	rightPairs := reversePairs1(nums, mid+1, right, tmp)

	if nums[mid] <= nums[mid+1] {
		return leftPairs + rightPairs
	}

	corePairs := mergeAndCount(nums, left, mid, right, tmp)
	return leftPairs + rightPairs + corePairs
}

func mergeAndCount(nums []int, left int, mid int, right int, tmp []int) int {
	for i := left; i <= right; i++ {
		tmp[i] = nums[i]
	}
	i := left
	j := mid + 1
	count := 0

	for k := left; k <= right; k++ {
		if i == mid+1 {
			nums[k] = tmp[i]
			j++
		} else if j == right+1 {
			nums[k] = tmp[i]
			i++
		} else if tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[i]
			j++
			count += mid - i + 1
		}
	}
	return count
}
