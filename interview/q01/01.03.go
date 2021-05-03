// Package q01
// @Description
// @Author 小游
// @Date 2021/05/01
package q01

import (
	"strings"
)

// 解法一，直接字符串替换
func replaceSpaces1(S string, length int) string {
	return strings.Replace(S[:length], " ", "%20", -1)
}

// 解法二，直接暴力
func replaceSpaces(S string, length int) string {
	// 先把S转换为字符数组
	data := []rune(S)
	// 使用两个指针
	i := length - 1
	j := len(S) - 1
	for i >= 0 {
		// 如果i位置为空格，那么我们就替换为02%
		if data[i] == ' ' {
			data[j] = '0'
			j--
			data[j] = '2'
			j--
			data[j] = '%'
			j--
		} else {
			data[j] = data[i]
			j--
		}
		i--
	}
	// 截取字符串
	return string(data[j+1:])
}
