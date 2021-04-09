// @Description
// @Author 小游
// @Date 2021/04/09
package q58_1

import (
	"bytes"
	"strings"
)

// 翻转单词顺序
func reverseWords1(s string) string {
	// 我们使用数组来暂存单词
	var words []string
	// 开始进行查找
	for i := 0; i < len(s); i++ {
		// 当不为空格时我们开始进行遍历
		if s[i] != ' ' {
			// 我们开始进行查找,直到遇到空格为止
			j := i
			for j < len(s) && s[j] != ' ' {
				j++
			}
			// 找到后我们把单词放入数组中
			words = append(words, s[i:j])
			i = j
		}
	}
	// 拼接字符串
	var ans bytes.Buffer
	for i := len(words) - 1; i >= 0; i-- {
		ans.Write([]byte(words[i] + " "))
	}
	var data = ans.String()
	// 去除最后的空格
	if len(data) != 0 {
		data = data[:len(data)-1]
	}
	return data
}

// 大佬的简化版本
func reverseWords(s string) string {
	// 首先去除空格
	s = strings.TrimSpace(s)
	// 我们使用两个指针来暂存数据
	i := len(s) - 1
	j := i
	// 使用bytes.Buffer加快拼接速度
	res := bytes.Buffer{}
	for i >= 0 {
		// 首先找到第一个空格
		for i >= 0 && s[i] != ' ' {
			i--
		}
		// 这个时候我们i和j之间就为一个单词
		res.Write([]byte(s[i+1:j+1] + " "))
		// 下面我们进行下一轮查找，我们这里跳过所有的空格
		for i >= 0 && s[i] == ' ' {
			i--
		}
		// 更新一下j的位置，用于下一轮查找
		j = i
	}
	return strings.TrimSpace(res.String())
}
