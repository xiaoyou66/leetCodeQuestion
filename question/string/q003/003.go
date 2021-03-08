// @Description
//	https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//	https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters/
// @Author 小游
// @Date 2021/03/07
package q003

// 解法一 我的暴力破解法
//func lengthOfLongestSubstring(s string) int {
//	i,j:= 0,0
//	max:=""
//	// 遍历s字符串
//	for i=0;i<len(s);i++ {
//		// 为了判断是否重复，这里我们使用一个map
//		m:=make(map[string]int)
//		// 第二层遍历，这里我们从i的位置开始
//		for j=i;j<len(s);j++ {
//			// 先判断是否重复
//			if _,ok:=m[string(s[j])];!ok{
//				// 不存在就插入
//				m[string(s[j])] = j
//			} else {
//				break
//			}
//			// 这里我们判断当前的长度是否为最长，如果是就返回
//			if j-i+1 > len(max) {
//				max = s[i:j+1]
//			}
//		}
//	}
//	return len(max)
//}

// 解法二 使用位图
func lengthOfLongestSubstring(s string) int {
	// 当s的长度为0时直接退出
	if len(s) == 0 {
		return 0
	}
	// 初始化位图
	var bitSet [256]bool
	// 初始化结果和左右两个指针
	result, left, right := 0, 0, 0
	// 当左指针指向s的右侧时，那么就退出循环
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将X标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		//
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}
