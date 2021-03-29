// @Description
// @Author 小游
// @Date 2021/03/27
package q227

import "unicode"

func calculate(s string) int {
	i := 0
	return parseExpr(s, &i)
}

func parseExpr(s string, i *int) int {
	op := "+"
	left, right := 0, 0
	// 使用i指针来记录当前的位置
	for *i < len(s) {
		// 首先确保我们的字符串不是空格
		if string(s[*i]) != " " {
			// 然后我们就解析数字，这我们i会解析到操作符所在的位置
			n := parseNum(s, i)
			// 下面判断操作类型，并进行不同的计算
			switch op {
			case "+":
				// 这里我们让right等于n，相当于暂存
				left += right
				right = n
			case "-":
				left += right
				right = -n
			case "*":
				// 如果是*的话优先级更高，优先计算
				right *= n
			case "/":
				right /= n
			}
			// 确保i不越界的情况下我们获取当前的操作符
			if *i < len(s) {
				op = string(s[*i])
			}
		}
		*i++
	}
	return left + right
}

// 解析数字
func parseNum(s string, i *int) int {
	n := 0
	// 我们这里不断一位一位进行解析，直到解析到操作符为止
	for *i < len(s) && unicode.IsDigit(rune(s[*i])) {
		n = 10*n + (int(s[*i]) - 48)
		*i++
	}
	return n
}
