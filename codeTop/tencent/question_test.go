// Package tencent
// @Description
// @Author 小游
// @Date 2021/04/15
package tencent

import (
	"fmt"
	"testing"
)

func Test_StringAdd(t *testing.T) {
	fmt.Println(addStrings("456", "77"))
}

var test = []int{1, 2, 3, 4, 5}

func Test_sun(t *testing.T) {
	sum := make([]int, len(test))
	sum[0] = test[0]
	for i := 1; i < len(test); i++ {
		sum[i] = sum[i-1] + test[i]
	}
	fmt.Println(sum[3] - sum[1])
}

func test1(arr []int) {
	i := 3
	j := 2
	arr[i] = arr[i] + arr[j]
	arr[j] = arr[i] - arr[j]
	arr[i] = arr[i] - arr[j]
	fmt.Println(arr[i])
}

func Test_Reverse(t *testing.T) {
	test1([]int{5, 8, 10, 20, 43})
	fmt.Println()
}
