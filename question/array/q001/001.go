// @Description https://leetcode-cn.com/problems/two-sum/submissions/
// @Author 小游
// @Date 2021/03/06
package q001

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
你可以按任意顺序返回答案。
示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func twoSum(nums []int,target int) []int {
	// 初始化一个map
	m:=make(map[int]int)
	// 遍历nums数组
	for i:=0;i<len(nums);i++{
		// 根据结果来获取我们需要的值
		answer:=target - nums[i]
		// 判断在结果在map中是否存在
		if _,ok:=m[answer];ok{
			return []int{m[answer],i}
		}
		// 存储我们的位置信息
		m[nums[i]] = i
	}
	return nil
}
