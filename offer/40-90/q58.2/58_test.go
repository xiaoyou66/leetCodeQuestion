// @Description
// @Author 小游
// @Date 2021/04/09
package q58_2

import (
	"bytes"
	"fmt"
	"testing"
)

func printTest(s string) string {
	// 判断第一个是数字还是字符串
	if s[0] > '0' && s[0] <= '9' {
		// 如果是数字，我们直接打印对应的次数
		return getCount(s[2:len(s)-1], int(s[0]-'0'))
	} else {
		// 如果是字母，我们只打印一次
		return s[:1] + getCount(s[1:], 1)
	}
}

func getCount(s string, n int) string {
	// 初始化buffer对象
	res := bytes.Buffer{}
	// 遍历n次
	for i := 0; i < n; i++ {
		// 判断第一个字符是数字还是字母
		if s[0] > '0' && s[0] <= '9' {
			// 如果是数字，我们按照同样的规则进行遍历
			res.Write([]byte(getCount(s[2:len(s)-1], int(s[0]-'0'))))
		} else {
			// 如果只剩下一个字符串，我们就只写一个
			if len(s) == 1 {
				res.Write([]byte(s[:1]))
			} else {
				// 否则我们继续遍历其余的
				res.Write([]byte(s[:1] + getCount(s[1:], 1)))
			}
		}
	}
	return res.String()
}

func Test_question(t *testing.T) {
	fmt.Println(printTest("2{A4{B5{C}}}"))
	fmt.Println(printTest("A4{B}"))
}
