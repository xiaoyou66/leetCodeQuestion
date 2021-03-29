// @Description
// @Author 小游
// @Date 2021/03/29
package q5

import (
	"strings"
)

// 第一个方法，使用自带的函数
func replaceSpace1(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// 第二个方法，一个一个遍历替换
func replaceSpace2(s string) string {
	// 遍历所有字符串
	for i := 0; i < len(s); i++ {
		// 如果字符串为空，那么我们就进行替换
		if string(s[i]) == " " {
			// 我们简单对字符串进行切片拼接
			s = s[:i] + "%20" + s[i+1:]
			//
			i += 2
		}
	}
	return s
}
