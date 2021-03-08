// @Description https://leetcode-cn.com/problems/3sum/
// @Author 小游
// @Date 2021/03/06
package q015

import "sort"

// 解法一 最优解，双指针 + 排序
//func threeSum(nums []int) [][]int {
//	// 对数据进行排序
//	sort.Ints(nums)
//	// 数值初始化
//	start,index,end,addNum,result,length := 0,0,0,0,make([][]int,0),len(nums)
//	// 遍历值
//	for index = 1;index < length-1;index++{
//		// 初始化start和end
//		start,end = 0,length-1
//		// 判断index 是否和上一步相同
//		if index > 1 && nums[index] == nums[index-1] {
//			start = index - 1
//		}
//		// 开始进行计算 这里需要同时满足条件才能执行
//		for start < index && end > index{
//			// 如果start和上一步相同，我们就跳过
//			if start > 0 && nums[start] == nums[start-1]{
//				start ++
//				continue
//			}
//			// 如果end和上一步相同，我们也跳过
//			if end < length - 1 && nums[end] == nums[end+1]{
//				end --
//				continue
//			}
//			// 正式开始判断
//			addNum  = nums[start] + nums[index] + nums[end]
//			if addNum == 0{
//				result = append(result, []int{nums[start],nums[index],nums[end]})
//				start ++
//				end --
//			} else if addNum > 0{
//				end --
//			} else {
//				start ++
//			}
//		}
//	}
//	return result
//}

// 解法二，使用map
func threeSum(nums []int) [][]int {
	// 结果数组
	res := [][]int{}
	counter:=make(map[int]int)
	for _,v:=range nums{
		counter[v] ++
	}
	uni:=[]int{}
	for k:=range counter{
		uni = append(uni,k)
	}
	sort.Ints(uni)
	for k,v:=range uni{
		if v*3 == 0 && counter[v] >=3 {
			res = append(res,[]int{v,v,v})
			continue
		}
		for k2:=k+1;k2<len(uni);k2++{
			v2 := uni[k2]
			// 如果出现两个相同的情况
			if v2*2+v==0 && counter[v2] > 1{
				res = append(res,[]int{v2,v2,v})
			}
			if v*2+v2==0 && counter[v] > 1{
				res = append(res,[]int{v2,v,v})
			}
			c:=0-v2-v
			if c > v2 && counter[c] > 0{
				res = append(res,[]int{v2,v,c})
			}
		}
	}
	return res
}
