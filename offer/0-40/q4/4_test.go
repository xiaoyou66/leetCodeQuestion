// @Description
// @Author 小游
// @Date 2021/03/29
package q4

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	fmt.Println(findNumberIn2DArray([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 5))
}
