// @Description
// @Author 小游
// @Date 2021/04/08
package q48

// 解法一 动态规划+哈希表
func lengthOfLongestSubstring1(s string) int {
	// 使用动态规划
	dic := make(map[byte]int)
	// res存储最终结果，然后tmp表示dp[i]
	res := 0
	tmp := 0
	// 遍历整个s
	for j := 0; j < len(s); j++ {
		// 默认i的值为-1，表示dic里面没有
		var i = -1
		// 如果当前字符串在map里面存在，我们就返回当前字符串的位置
		if _, ok := dic[s[j]]; ok {
			i = dic[s[j]]
		}
		// 更新一下map
		dic[s[j]] = j
		// 判断是当前的大还是j-i的大
		if tmp < j-i {
			tmp++
		} else {
			tmp = j - i
		}
		// 获取最大值
		res = max(res, tmp)
	}
	return res
}

// 解法二 动态规划+线性表
func lengthOfLongestSubstring2(s string) int {
	// res表示最终结果，tmp表示当前位置下最长的子字符串
	res := 0
	tmp := 0
	// 开始进行遍历
	for i := 0; i < len(s); i++ {
		j := i - 1
		// 找出第一个和当前位置重复的元素
		for j >= 0 && s[j] != s[i] {
			j--
		}

		if tmp < i-j {
			tmp++
		} else {
			tmp = i - j
		}
		// 找出一个最大值
		res = max(res, tmp)
	}
	return res
}

// 解法三 双指针+哈希表
func lengthOfLongestSubstring(s string) int {
	// 使用dic来暂存数据
	dic := make(map[byte]int)
	// res设置为0，j默认为-1
	res := 0
	j := -1
	// 遍历整个字符串
	for i := 0; i < len(s); i++ {
		// 当我们的字符串已经存在时
		if _, ok := dic[s[i]]; ok {
			// 这里我们更新一下j的位置（取最大是为了每次获取最后的，避免重复）
			j = max(j, dic[s[i]])
		}
		// 然后我们更新一下位置信息
		dic[s[i]] = i
		// 更新一下最大值即可
		res = max(res, i-j)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
