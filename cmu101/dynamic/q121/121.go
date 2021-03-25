// @Description https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// @Author 小游
// @Date 2021/03/25
package q121

import "math"

func maxProfit(prices []int) int {
	// sell表示利润
	sell := 0
	// buy表示我们买入的价格，这里我们取负，后面计算利润可以直接相加即可
	buy := math.MinInt32
	// 遍历整个数组
	for _, v := range prices {
		// 首先找到一个最低价格买入
		buy = max(buy, -v)
		// 判断当前卖出是否可以获得最大利润
		sell = max(sell, buy+v)
	}
	// 返回最大利润
	return sell
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
