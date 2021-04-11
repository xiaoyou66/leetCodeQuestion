// @Description
// @Author 小游
// @Date 2021/04/11
package important

import (
	"fmt"
	"testing"
)

// 最长公共子串测试
func Test_LCS(t *testing.T) {
	fmt.Println(getLCS2("ABCDEFG", "ABZDEFKG"))
}

func Test_coin(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
