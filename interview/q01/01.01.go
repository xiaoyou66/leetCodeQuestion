// Package q01
// @Description
// @Author 小游
// @Date 2021/04/14
package q01

// 最简单的解法
func isUnique1(astr string) bool {
	// 使用map暂存
	unique := make(map[int32]bool)
	for _, v := range astr {
		if unique[v] {
			return false
		}
		unique[v] = true
	}
	return true
}

// 方法二 布尔数组
func isUnique2(astr string) bool {
	unique := make([]bool, 26)
	for i := 0; i < len(astr); i++ {
		if unique[astr[i]-'a'] {
			return false
		}
		unique[astr[i]-'a'] = true
	}
	return true
}

// 方法三 位运算
func isUnique(astr string) bool {
	mask := 0
	for i := 0; i < len(astr); i++ {
		bitMove := astr[i] - 'a'
		if mask&(1<<bitMove) != 0 {
			return false
		} else {
			mask |= 1 << bitMove
		}
	}
	return true
}
