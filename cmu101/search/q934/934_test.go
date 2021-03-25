// @Description
// @Author 小游
// @Date 2021/03/17
package q934

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	fmt.Println(shortestBridge([][]int{{0, 1}, {1, 0}}))
	fmt.Println(shortestBridge([][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}))
}
