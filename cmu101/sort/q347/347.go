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
	var ans []int
	for i:=len(bucket)-1;i>=len(bucket)-k;i-- {
		ans = append(ans,bucket[i]...)
	}
	return ans
}