// @Description  https://leetcode-cn.com/problems/candy/
// @Author 小游
// @Date 2021/03/11
package q135

func candy(ratings []int) int {
	// 初始化数组并全部赋值为1
	ans:=make([]int,len(ratings))
	for k:=range ans{
		ans[k] = 1
	}
	// 首先我们先满足评分更高的比左边的糖果更多
	for i:=1;i<len(ratings);i++{
		if ratings[i] > ratings[i-1] {
			// 如果评分更高，那么当前孩子要比左边的孩子糖果要多1
			ans[i] = ans[i-1] + 1
		}
	}
	// 然后我们满足评分更高的比右边的糖果更多，这里我们就从右边开始遍历
	for i:=len(ratings)-1;i>0;i--{
		// 如果当前孩子比左边的孩子评分低，那么就修改左边孩子的值
		if ratings[i] < ratings[i-1]  {
			// 左边孩子的评分比右边的高，且左边孩子当前的糖果数
			// 不大于右边孩子的糖果数，则左边孩子的糖果数更新为右边孩子的糖果数加 1
			if ans[i]+1 > ans[i-1] {
				ans[i-1] = ans[i] + 1
			}
		}
	}
	a:=0
	// 计算糖果数
	for _,v:=range ans {
		a+=v
	}
	return a
}