// Package q01
// @Description
// @Author 小游
// @Date 2021/04/14
package q01

import (
	"sort"
	"strings"
)

// CheckPermutation1 方法一字符统计
func CheckPermutation1(s1 string, s2 string) bool {
	// 使用data来存储数据
	data := make([]int, 26)
	// 首先遍历s1，统计一下
	for _, v := range s1 {
		data[v-'a']++
	}
	// 然后遍历s2，每个字符减1
	for _, v := range s2 {
		data[v-'a']--
	}
	// 最后我们遍历数据表，如果不为0直接返回false
	for _, v := range data {
		if v != 0 {
			return false
		}
	}
	return true
}

// 对字符串进行排序
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	tmp1 := strings.Split(s1, "")
	tmp2 := strings.Split(s2, "")
	sort.Strings(tmp1)
	sort.Strings(tmp2)
	return strings.Join(tmp1, "") == strings.Join(tmp2, "")
}
