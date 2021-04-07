// @Description
// @Author 小游
// @Date 2021/03/30
package q12

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
}
