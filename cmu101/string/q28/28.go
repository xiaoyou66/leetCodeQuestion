// @Description
// @Author 小游
// @Date 2021/03/27
package q28

func strStr(haystack string, needle string) int {
	k, n, p := -1, len(haystack), len(needle)
	// 如果need为0，那么我们直接返回0
	if p == 0 {
		return 0
	}
	// 初始化next数组
	next := make([]int, p)
	for i := 0; i < p; i++ {
		next[i] = -1
	}
	// 计算next数组
	calNext(needle, next)
	// 开始进行字符串查找
	for i := 0; i < n; i++ {
		// k表示needle的位置，如果为-1，那么我们就先不计算
		for k > -1 && needle[k+1] != haystack[i] {
			k = next[k]
		}
		// 这里我们
		if needle[k+1] == haystack[i] {
			k++
		}
		if k == p-1 {
			return i - p + 1
		}
	}
	return -1
}

// 计算next数组
// 简单介绍一next数组
// ababa 生成的是 -1 -1 0 1 2
func calNext(needle string, next []int) {
	// 默认情况下，我们的next数组下标为-1（第一个一定是-1）
	p := -1
	// 然后我们从1开始往后面进行计算
	for j := 1; j < len(needle); j++ {
		// 这里我们使用回溯来计算next数组
		for p > -1 && needle[p+1] != needle[j] {
			p = next[p]
		}
		// 当有一位相同时，p要不断++
		if needle[p+1] == needle[j] {
			p++
		}
		next[j] = p
	}
}
