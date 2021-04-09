// @Description
// @Author 小游
// @Date 2021/04/09
package q58_2

import "bytes"

// 解法一 go字符串拼接
func reverseLeftWords1(s string, n int) string {
	return s[n:] + s[:n]
}

// 解法二： 列表遍历拼接
func reverseLeftWords(s string, n int) string {
	buffer := bytes.Buffer{}
	for i := n; i < n+len(s); i++ {
		buffer.Write([]byte{s[i%len(s)]})
	}
	return buffer.String()
}
