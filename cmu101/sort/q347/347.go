// @Description 
// @Author 小游
// @Date 2021/03/15
package q347

func topKFrequent(nums []int, k int) []int {
	// 首先我们使用count来统计出现的频次
	count:=make(map[int]int)
	// 先遍历一次获取出现频率最大的
	max:=0
	for _,v:=range nums{
		count[v] ++
		// 这里获取出现次数最大的值
		if count[v] > max {
			max = count[v]
		}
	}
	// 然后我们就可以按照出现的频率来创建一个桶
	bucket:=make([][]int,max+1)
	// 创建一个map来存储数据
	for i,v:=range count{
		bucket[v]=append(bucket[v], i)
	}
	// 下面我们就可以找出前n高的元素了
	var ans []int
	for i:=len(bucket)-1;i>=0;i-- {
		// 因为我们桶里面是可能出现不同数的情况所以我们需要遍历
		for _,v:=range bucket[i] {
			// 当我们找出这前n个时，我们就立即返回
			if len(ans) < k {
				ans = append(ans,v)
			} else {
				return ans
			}
		}
	}
	return ans
}