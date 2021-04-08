// @Description
// @Author 小游
// @Date 2021/04/08
package q50

func firstUniqChar(s string) byte {
	cnt := [26]int{}
	// 遍历整个数组
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	// 继续遍历
	for i := 0; i < len(s); i++ {
		// 找出第一个为1的值，找到后我们立马退出
		if cnt[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
