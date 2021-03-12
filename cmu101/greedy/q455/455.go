// @Description 分发饼干 https://leetcode-cn.com/problems/assign-cookies/
// @Author 小游
// @Date 2021/03/11
package q455

import "sort"

// 解题思路：就是我们要优先用最小的饼干来满足最小的要求
// 所以我们可以通过排序，让最小的孩子和最小的饼干进行测试，然后我们优先满足最小需求的孩子
func findContentChildren(g []int, s []int) int {
	// 首先我们给g和s进行排序
	sort.Ints(g)
	sort.Ints(s)
	// 分别用i，j来表示孩子和饼干
	i,j:=0,0
	// 当任意一个不满足时，我们就可以退出了
	for i<len(g) && j <len(s) {
		// 判断当前饼干是否满足当前孩子，如果满足孩子+1
		if g[i] <= s[j] {
			i ++
		}
		// 无论是否满足，饼干都要后移
		j++
	}
	return i
}