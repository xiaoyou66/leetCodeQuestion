// Package tencent
// @Description
// @Author 小游
// @Date 2021/04/18
package tencent

func addStrings(num1 string, num2 string) string {
	// 为了方便计算，我们直接把字符串转换为byte
	str1 := []byte(num1)
	str2 := []byte(num2)
	l1 := len(str1) - 1
	l2 := len(str2) - 1
	// 同理我们获取两个字符串的长度，然后获取最大的一个来新建一个byte数组
	// 这个byte数组用于存储结果
	high := max(l2, l1)
	res := make([]byte, high+1)
	// add表示进位
	add := 0
	// 当两个字符串指针都加完时我们跳出循环
	for l1 >= 0 || l2 >= 0 {
		sum := 0
		// 因为可能存在一个指针已经加完的情况，所以我们这里为了三种情况来计算
		// 注意这里我们需要考虑进位。我们使用add来表示
		if l1 >= 0 && l2 >= 0 {
			sum = int(str1[l1]-'0'+str2[l2]-'0') + add
		} else if l1 >= 0 && l2 < 0 {
			sum = int(str1[l1]-'0') + add
		} else {
			sum = int(str2[l2]-'0') + add
		}
		// 如果sum大于0，表示有进位，否则我们把add置为0
		if sum >= 10 {
			add = sum / 10
		} else {
			add = 0
		}
		// 这里我们就可以知道当前结果位结果是多少了
		res[high] = byte(sum%10) + '0'
		high--
		l1--
		l2--
	}
	// 如果还存在进位的话，我们就需要把进位加进去
	if add != 0 {
		return string(rune(add+'0')) + string(res)
	} else {
		return string(res)
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
