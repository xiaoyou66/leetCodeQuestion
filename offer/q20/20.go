// @Description
// @Author 小游
// @Date 2021/04/04
package q20

func isNumber(s string) bool {
	// 状态值
	states := []map[byte]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4},
		{'d': 2, '.': 4},
		{'d': 2, '.': 3, 'e': 5, ' ': 8},
		{'d': 3, 'e': 5, ' ': 8},
		{'d': 3},
		{'s': 6, 'd': 7},
		{'d': 7},
		{'d': 7, ' ': 8},
		{' ': 8},
	}
	p := 0
	var t byte
	for _, c := range s {
		// 这里是判断当前字符串是属于什么类型
		if c >= '0' && c <= '9' {
			t = 'd'
		} else if c == '+' || c == '-' {
			t = 's'
		} else if c == 'e' || c == 'E' {
			t = 'e'
		} else if c == '.' || c == ' ' {
			t = byte(c)
		} else {
			t = '?'
		}
		// 如果当前状态不存在，那么就返回false
		if _, ok := states[p][t]; !ok {
			return false
		}
		// 进行状态转移
		p = states[p][t]
	}
	return p == 2 || p == 3 || p == 7 || p == 8
}
