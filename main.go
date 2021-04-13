// @Description
// @Author 小游
// @Date 2021/03/06
package main

/**
为了寻找最佳拍档，我们定义两人名字的缘分值：两人名字左对齐后，对应位置字的拼音的缘分值之和。
对于两个拼音s1、s2，通过剔除一些字符使得留下的子串一模一样，被剔除字符之和的最小值即为两个拼音的缘分值
。求给定两人名字的缘分值。

字符之和是指对应的ASCII值之和。
样例中，Zhang和Zhan去掉‘g’后，剩余子串相等，因此第一个字的缘分值为asc(‘g’)=103；
而San和Ai则需全部删除San和Ai，其缘分值为asc(“San”)+asc(“Ai”)=290+170=460。
因此Zhang San和Zhan Ai的缘分值为103+460=563。
边界场景，一人名字较长者，譬如，Ali Ba Ba和Xie Cheng，那最后一个Ba需要全部剔除掉。
*/

func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	var dummy [3]int
	for i := 0; i < len(dummy); i++ {
		println(i) // 0, 1, 2
	}
}
