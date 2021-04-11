// @Description
// @Author 小游
// @Date 2021/04/10
package q63

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profile := 0
	sell := prices[0]
	for i := 0; i < len(prices); i++ {
		// 判断当前价格是否比sell还低
		if prices[i] < sell {
			sell = prices[i]
		} else if prices[i]-sell > profile {
			profile = prices[i] - sell
		}
	}
	return profile
}
