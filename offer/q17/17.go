// @Description
// @Author 小游
// @Date 2021/03/31
package q17

import (
	"math"
)

// 解法一 暴力
func printNumbers1(n int) []int {
	// 先获取对应容量
	target := int(math.Pow(10, float64(n)))
	data := make([]int, target-1)
	for i := 0; i < target-1; i++ {
		data[i] = i + 1
	}
	return data
}

// 解法二 分治递归法
var nine = 0
var start, n1 int
var num []byte
var loop = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var res []byte

func printNumbers(n int) string {
	n1 = n
	// 为了避免字符串出现0的问题，我们需要使用start来进行截取
	start = n - 1
	// 初始化num数组，用于存放数据
	num = make([]byte, n)
	// 从第0位开始进行遍历
	dfs(0)
	// 最后返回的时候记得把逗号去掉
	return string(res[:len(res)-1])
}

func dfs(x int) {
	// 当x等于n1的时候我们就打印完毕，可以跳出循环了
	if x == n1 {
		// 转换的时候注意字符串截取
		s := string(num[start:])
		// 当我们的s不为0时才把数据加进去
		if s != "0" {
			res = append(res, s...)
			res = append(res, ',')
		}
		// 这里表示我们的数字进了一位
		if n1-start == nine {
			start--
		}
		return
	}
	// 遍历一下loop数组
	for _, v := range loop {
		// 当v为9时，我们就需要进一位
		if v == '9' {
			nine++
		}
		// 设置当前位数（因为前面几位都是固定的，所以我们只需要设置x位）
		num[x] = v
		// 计算下一位
		// 这里我们需要遍历9次
		dfs(x + 1)
	}
	nine--
}
