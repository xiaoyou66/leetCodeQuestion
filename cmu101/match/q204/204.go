// @Description
// @Author 小游
// @Date 2021/03/25
package q204

import "math"

// 方法一 埃拉托斯特尼筛法
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	prime := make([]bool, n+1)
	count := n / 2 // 偶数一定不是质数
	i := 3
	sqrt := int(math.Sqrt(float64(n)))
	// 最小质因子一定小于等于开方数
	for i <= sqrt {
		// 如果这个数没有标记，那么我们就进行计算
		if !prime[i] {
			// 这里我们直接计算以i为倍数时，把所有小于n且是i的倍速的全部标记一下
			for j := i * i; j < n; j += 2 * i {
				// 如果当前数据没有标记，那么我们就标记一下，并让count-1
				if !prime[j] {
					prime[j] = true
					count--
				}
			}
			for {
				// 这里我们不断遍历，把偶数去掉
				i += 2
				if !(i <= sqrt && prime[i]) {
					break
				}
			}
		}
	}
	// 最后剩下的count其实就是质数了
	return count
}
