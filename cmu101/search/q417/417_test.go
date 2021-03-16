// @Description
// @Author 小游
// @Date 2021/03/16
package q417

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	fmt.Println(pacificAtlantic([][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}))
}
