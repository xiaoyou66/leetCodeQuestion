// @Description
// @Author 小游
// @Date 2021/03/26
package q242

func isAnagram(s string, t string) bool {
	// 如果长度不一样，那么直接退出
	if len(s) != len(t) {
		return false
	}
	// 因为只有小写字母，所以我们可以只创建一个大小为26的数组
	a := make([]int, 26)
	for k := range s {
		// 如果两个字符串相等的话，那么一个++一个--后，整个数组应该是为0的
		a[s[k]-'a']++
		a[t[k]-'a']--
	}
	// 最后只要发现一个字符串不为0，那么我们就直接返回false
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}
