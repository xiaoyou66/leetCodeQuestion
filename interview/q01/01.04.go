// Package q01 
// @Description 
// @Author 小游
// @Date 2021/05/03
package q01

func canPermutePalindrome1(s string) bool {
	// 创建一个map来暂存数据
	tmp:=make(map[rune]bool)
	// 遍历字符串
	for	_,c:=range s{
		// 判断字符串是否存在
		if tmp[c] {
			delete(tmp,c)
		} else {
			tmp[c] = true
		}
	}
	return len(tmp) <= 1
}

func canPermutePalindrome(s string) bool {
	// 我们使用result数组来存储最终结果
	var result [4]int
	// 遍历字符串s来把我们的值放入到这个数组中
	for _, v := range s {
		atoi := int(v)
		// 这里的作用就是使用异或运算，相同置0，不同置1
		result[atoi/32] ^= 1 << (atoi % 32)
	}

	var cnt int

	for i := 0; i < 4; i ++ {
		// 判断result[i]这部分位数是否只有一个bit为1
		if (result[i] & -result[i]) == result[i] {
			if result[i] != 0 {
				cnt ++
			}
		} else { // 如果这段bit不止一个1, 直接判断false
			return false
		}
	}

	// 只有一个置位或者没有则是回文字符串
	return cnt <= 1
}