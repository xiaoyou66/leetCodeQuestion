// @Description
// @Author 小游
// @Date 2021/03/30
package q9

import (
	"fmt"
	"testing"
)

func Test_question(t *testing.T) {
	obj := Constructor()
	obj.AppendTail(4)
	obj.AppendTail(5)
	param2 := obj.DeleteHead()
	fmt.Println(param2)
}
